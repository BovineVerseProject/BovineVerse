// SPDX-License-Identifier: MIT
pragma solidity 0.8.4;

import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";
interface IBEP20 {
    function decimals() external view returns (uint8);

    function name() external view returns (string memory);

    function symbol() external view returns (string memory);

    function totalSupply() external view returns (uint);

    function balanceOf(address account) external view returns (uint);

    function transfer(address recipient, uint amount) external returns (bool);

    function allowance(address owner, address spender) external view returns (uint);

    function approve(address spender, uint amount) external returns (bool);

    function transferFrom(address sender, address recipient, uint amount) external returns (bool);

    function mint(address addr_, uint amount_) external;

    event Transfer(address indexed from, address indexed to, uint value);
    event Approval(address indexed owner, address indexed spender, uint value);
}

contract BVGMining is ReentrancyGuard {
    IBEP20 public BVG = IBEP20(0x96cB0Ade2e254c598b12179503C00b007EeB7861);
    using SafeERC20 for IERC20;
    uint public startTime;
    uint constant miningTime = 30 days;
    uint public totalClaimed;
    uint constant dayliyOut = 5000000 ether;
    uint public rate;
    uint constant acc = 1e10;
    address public owner;
    constructor (){
        rate = dayliyOut / 86400;
        poolInfo[1].rate = 50;
        poolInfo[2].rate = 30;
        poolInfo[3].rate = 20;
        poolInfo[1].limit = 10 ether;
        poolInfo[2].limit = 5000 ether;
        poolInfo[3].limit = 100 ether;
        poolInfo[2].token = 0x55d398326f99059fF775485246999027B3197955;
        //USDT
        poolInfo[3].token = 0x0D8Ce2A99Bb6e3B7Db580eD848240e4a0F9aE153;
        //FIL
        owner = msg.sender;
    }
    struct PoolInfo {
        address token;
        uint TVL;
        uint lastTime;
        uint debt;
        uint claimed;
        uint rate;
        uint limit;
        uint lastDebt;
    }

    struct UserInfo {
        uint stakeAmount;
        uint debt;
        uint toClaim;
    }

    mapping(uint => PoolInfo) public poolInfo;
    mapping(address => mapping(uint => UserInfo)) public userInfo;
    mapping(address => uint) public userClaimed;

    event Claim(address indexed player, uint indexed amount);
    event Stake(address indexed player, uint indexed amount, uint indexed poolId);
    event UnStake(address indexed player, uint poolID);
    
    modifier checkEnd(){
        if(block.timestamp >= startTime + miningTime && poolInfo[1].lastDebt ==0){
            poolInfo[1].lastDebt = coutingDebt(1);
            poolInfo[2].lastDebt = coutingDebt(2);
            poolInfo[3].lastDebt = coutingDebt(3);
        }
        _;
    }

    function coutingDebt(uint poolId) public view returns (uint _debt){
        PoolInfo storage info = poolInfo[poolId];
        if (info.lastDebt != 0){
            return info.lastDebt;
        }
        _debt = info.TVL > 0 ? (rate * info.rate / 100) * (block.timestamp - info.lastTime) * acc / info.TVL + info.debt : 0 + info.debt;
    }

    function calculateReward(address player, uint poolId) public view returns (uint){
        UserInfo storage user = userInfo[player][poolId];
        if (user.stakeAmount == 0) {
            return 0;
        }
        uint rew = user.stakeAmount * (coutingDebt(poolId) - user.debt) / acc;
        return (rew + user.toClaim);
    }

    function calculateAllReward(address player) public view returns (uint){
        uint rew;
        for (uint i = 1; i <= 3; i ++) {
            UserInfo storage user = userInfo[player][i];
            if (user.stakeAmount == 0) {
                continue;
            }
            rew += calculateReward(player, i);
        }
        return rew;
    }
    
    function setStartTime(uint time_) external{
        require(msg.sender == owner,'not owner');
        require(startTime == 0,'starting');
        require(block.timestamp < time_ + miningTime ,'out of time');
        require(time_ != 0 ,'startTime can not be zero');
        startTime = time_;
        owner = address(0);
    }

    function claimReward(uint poolId) internal checkEnd {
        uint rew = calculateReward(msg.sender, poolId);
        if (rew == 0) {
            return;
        }
        UserInfo storage user = userInfo[msg.sender][poolId];
        BVG.mint(msg.sender, rew);
        uint tempDebt = coutingDebt(poolId);
        user.debt = tempDebt;
        user.toClaim = 0;
        userClaimed[msg.sender] += rew;
        totalClaimed += rew;
        emit Claim(msg.sender, rew);

    }

    function claimAllReward() external checkEnd{
        // require(block.timestamp < startTime + miningTime, 'mining over');
        uint rew;
        uint tempDebt;
        for (uint i = 1; i <= 3; i ++) {
            UserInfo storage user = userInfo[msg.sender][i];
            if (user.stakeAmount == 0) {
                continue;
            }
            rew += calculateReward(msg.sender, i);
            tempDebt = coutingDebt(i);
            user.debt = tempDebt;
            user.toClaim = 0;
        }
        require(rew > 0, 'no reward');
        userClaimed[msg.sender] += rew;
        totalClaimed += rew;
        BVG.mint(msg.sender, rew);
        emit Claim(msg.sender, rew);
    }

    function stakeBnb() payable external {
        require(block.timestamp >= startTime && startTime != 0,'not start');
        require(block.timestamp < startTime + miningTime, 'mining over');
        require(msg.value > 0, 'amount can not be zero');
        UserInfo storage user = userInfo[msg.sender][1];
        PoolInfo storage pool = poolInfo[1]; 
        require(user.stakeAmount + msg.value <= pool.limit, 'out of limit');
        if (user.stakeAmount > 0) {
            user.toClaim = calculateReward(msg.sender, 1);
        }
        uint tempDebt = coutingDebt(1);
        pool.TVL += msg.value;
        pool.lastTime = block.timestamp;
        pool.debt = tempDebt;
        user.stakeAmount += msg.value;
        user.debt = tempDebt;
        emit Stake(msg.sender,msg.value,1);
    }

    function unStakeBnb() external nonReentrant checkEnd{
        UserInfo storage user = userInfo[msg.sender][1];
        PoolInfo storage pool = poolInfo[1];
        require(user.stakeAmount > 0, 'no stakeAmount');
        claimReward(1);
        uint tempDebt = coutingDebt(1);
        uint amount = user.stakeAmount;
        pool.TVL -= amount;
        pool.lastTime = block.timestamp;
        pool.debt = tempDebt;
        user.stakeAmount = 0;
        user.debt = tempDebt;
        payable(msg.sender).transfer(amount);
        emit UnStake(msg.sender,amount);
    }

    function stakeToken(uint poolId, uint amount) external  {
        require(block.timestamp >= startTime && startTime != 0,'not start');
        require(poolId == 2 || poolId == 3, 'wrong pool ID');
        require(block.timestamp < startTime + miningTime, 'mining over');
        require(amount > 0, 'amount can not be zero');
        UserInfo storage user = userInfo[msg.sender][poolId];
        PoolInfo storage pool = poolInfo[poolId];
        require(user.stakeAmount + amount <= pool.limit, 'out of limit');
        if (user.stakeAmount > 0) {
            user.toClaim = calculateReward(msg.sender, poolId);
        }
        uint tempDebt = coutingDebt(poolId);
        pool.TVL += amount;
        pool.lastTime = block.timestamp;
        pool.debt = tempDebt;
        user.stakeAmount += amount;
        user.debt = tempDebt;
        IERC20(pool.token).safeTransferFrom(msg.sender, address(this), amount);
        emit Stake(msg.sender,amount,poolId);
    }

    function unStakeToken(uint poolId) external nonReentrant checkEnd{
        require(poolId == 2 || poolId == 3, 'wrong pool ID');
        UserInfo storage user = userInfo[msg.sender][poolId];
        PoolInfo storage pool = poolInfo[poolId];
        require(user.stakeAmount > 0, 'no stakeAmount');
        claimReward(poolId);
        uint tempDebt = coutingDebt(poolId);
        uint amount = user.stakeAmount;
        pool.TVL -= amount;
        pool.lastTime = block.timestamp;
        pool.debt = tempDebt;
        user.stakeAmount = 0;
        user.debt = tempDebt;
        IERC20(pool.token).safeTransfer(msg.sender, amount);
        emit UnStake(msg.sender,amount);
    }


}