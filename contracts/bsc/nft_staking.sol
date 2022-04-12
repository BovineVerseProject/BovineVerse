// SPDX-License-Identifier: MIT
pragma solidity 0.8.4;

import "../interface/IERC_721.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC721/utils/ERC721HolderUpgradeable.sol";
import "../interface/IBVG.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

contract NFTStaking is ERC721HolderUpgradeable, OwnableUpgradeable {
    IBEP20 public BVG;
    uint public startTime;
    uint constant miningTime = 30 days;
    uint public totalClaimed;
    uint constant totalSupply = 50000000 ether;
    uint public rate;
    uint constant acc = 1e10;
    I721 public OAT;
    I721 public IGO;
    I721 public info;
    mapping(address => mapping(uint => address)) public cardOwner;
    uint public IGOPower;
    mapping(uint => uint) public OATPower;

    function initialize() public initializer {
        __Context_init_unchained();
        __Ownable_init_unchained();
        rate = totalSupply / miningTime;
        OATPower[1658] = 585;
        OATPower[1473] = 375;
        OATPower[1423] = 846;
        OATPower[1300] = 1270;
        OATPower[1196] = 274;
        OAT = I721(0x9F471abCddc810E561873b35b8aad7d78e21a48e);
        IGO = I721(0x927a7f35587BC7E59991CCCdd2D4aDd1f0e7bc66);
        info = I721(0xaf84c52D2117dADBD22FC440825e901E8d4E3BD2);
        BVG = IBEP20(0x96cB0Ade2e254c598b12179503C00b007EeB7861);
        IGOPower = 12700;
    }

    struct UserInfo {
        uint power;
        uint debt;
        uint toClaim;
        uint claimed;
        uint[] IGOList;
        uint[] OATList;
    }

    mapping(address => UserInfo) public userInfo;
    uint public debt;
    uint public totalPower;
    uint public lastTime;
    uint public lastDebt;


    mapping(uint => bool)public isOld;
    uint public ssss;
    event Claim(address indexed player, uint indexed amount);
    event Stake(address indexed player, uint indexed tokenId);
    event UnStake(address indexed player, uint indexed tokenId);

    modifier checkEnd(){
        if (block.timestamp >= startTime + miningTime && lastDebt == 0) {
            lastDebt = coutingDebt();
        }
        _;
    }
    function calculateReward(address player) public view returns (uint){
        UserInfo storage user = userInfo[player];
        if (user.power == 0 && user.toClaim == 0) {
            return 0;
        }
        uint rew = user.power * (coutingDebt() - user.debt) / acc;
        return (rew + user.toClaim);
    }

    function setStartTime(uint time_) external onlyOwner {
        require(startTime == 0, 'starting');
        require(block.timestamp < time_ + miningTime, 'out of time');
        require(time_ != 0, 'startTime can not be zero');
        startTime = time_;
    }

    function newOat(uint cid_, uint power, bool b) external onlyOwner {
        OATPower[cid_] = power;
        isOld[cid_] = b;
    }

    function coutingDebt() public view returns (uint _debt){
        if (lastDebt != 0) {
            return lastDebt;
        }
        _debt = totalPower > 0 ? rate * (block.timestamp - lastTime) * acc / totalPower + debt : 0 + debt;
    }

    function stakeIGO(uint tokenId) external {
        require(block.timestamp >= startTime && startTime != 0, 'not start');
        require(cardOwner[address(IGO)][tokenId] == address(0), 'staked');
        require(block.timestamp < startTime + miningTime, 'mining over');
        IGO.safeTransferFrom(msg.sender, address(this), tokenId);
        cardOwner[address(IGO)][tokenId] = msg.sender;
        if (userInfo[msg.sender].power > 0) {
            userInfo[msg.sender].toClaim = calculateReward(msg.sender);
        }
        userInfo[msg.sender].IGOList.push(tokenId);
        uint tempDebt = coutingDebt();
        userInfo[msg.sender].debt = tempDebt;
        userInfo[msg.sender].power += IGOPower;
        debt = tempDebt;
        totalPower += IGOPower;
        lastTime = block.timestamp;
        emit Stake(msg.sender, tokenId);
    }

    function stakeOAT(uint tokenId) external {
        require(block.timestamp >= startTime && startTime != 0, 'not start');
        require(cardOwner[address(OAT)][tokenId] == address(0), 'staked');
        require(block.timestamp < startTime + miningTime, 'mining over');
        OAT.safeTransferFrom(msg.sender, address(this), tokenId);
        cardOwner[address(OAT)][tokenId] = msg.sender;
        if (userInfo[msg.sender].power > 0) {
            userInfo[msg.sender].toClaim = calculateReward(msg.sender);
        }
        uint _cid = info.cid(tokenId);
        if (_cid == 0) {
            _cid = OAT.cid(tokenId);
        }
        uint tempPower = OATPower[_cid];
        require(tempPower > 0, 'wrong cid');
        userInfo[msg.sender].OATList.push(tokenId);
        uint tempDebt = coutingDebt();
        userInfo[msg.sender].debt = tempDebt;
        userInfo[msg.sender].power += tempPower;
        debt = tempDebt;
        totalPower += tempPower;
        lastTime = block.timestamp;
        emit Stake(msg.sender, tokenId);
    }

    function unStakeIGO(uint tokenId) external {
        require(cardOwner[address(IGO)][tokenId] == msg.sender, 'not card owner');
        delete cardOwner[address(IGO)][tokenId];
        uint tempRew = calculateReward(msg.sender);
        if (tempRew > 0) {
            userInfo[msg.sender].toClaim = tempRew;
        }
        uint tempDebt = coutingDebt();
        userInfo[msg.sender].debt = tempDebt;
        userInfo[msg.sender].power -= IGOPower;
        debt = tempDebt;
        totalPower -= IGOPower;
        lastTime = block.timestamp;
        uint index;
        uint length = userInfo[msg.sender].IGOList.length;
        for (uint i = 0; i < length; i ++) {
            if (userInfo[msg.sender].IGOList[i] == tokenId) {
                index = i;
                break;
            }
        }
        userInfo[msg.sender].IGOList[index] = userInfo[msg.sender].IGOList[length - 1];
        userInfo[msg.sender].IGOList.pop();
        IGO.safeTransferFrom(address(this), msg.sender, tokenId);
        emit UnStake(msg.sender, tokenId);
    }

    function unStakeOAT(uint tokenId) external {
        require(cardOwner[address(OAT)][tokenId] == msg.sender, 'not card owner');
        delete cardOwner[address(OAT)][tokenId];
        uint tempRew = calculateReward(msg.sender);
        if (tempRew > 0) {
            userInfo[msg.sender].toClaim = tempRew;
        }
        uint tempDebt = coutingDebt();
        uint _cid = info.cid(tokenId);
        if (_cid == 0) {
            _cid = OAT.cid(tokenId);
        }
        uint tempPower = OATPower[_cid];
        userInfo[msg.sender].debt = tempDebt;
        userInfo[msg.sender].power -= tempPower;
        debt = tempDebt;
        totalPower -= tempPower;
        lastTime = block.timestamp;
        uint index;
        uint length = userInfo[msg.sender].OATList.length;
        for (uint i = 0; i < length; i ++) {
            if (userInfo[msg.sender].OATList[i] == tokenId) {
                index = i;
                break;
            }
        }
        userInfo[msg.sender].OATList[index] = userInfo[msg.sender].OATList[length - 1];
        userInfo[msg.sender].OATList.pop();
        OAT.safeTransferFrom(address(this), msg.sender, tokenId);
        emit UnStake(msg.sender, tokenId);
    }

    function claimReward() external {
        uint rew;
        rew = calculateReward(msg.sender);
        require(rew > 0, 'no reward');
        userInfo[msg.sender].claimed += rew;
        userInfo[msg.sender].debt = coutingDebt();
        userInfo[msg.sender].toClaim = 0;
        totalClaimed += rew;
        BVG.mint(msg.sender, rew);
        emit Claim(msg.sender, rew);
    }

    function checkUserOATList(address addr) public view returns (uint[] memory, uint[] memory){
        uint[] memory list = new uint[](userInfo[addr].OATList.length);

        for (uint i = 0; i < list.length; i ++) {
            list[i] = info.cid(userInfo[addr].OATList[i]);
            if(list[i] == 0){
                list[i] = OAT.cid(userInfo[addr].OATList[i]);
            }
        }
        return (userInfo[addr].OATList, list);
    }

    function checkUserIGOList(address addr) public view returns (uint[] memory){
        return userInfo[addr].IGOList;
    }

    function withDrawToken(address token_, address wallet, uint amount) external onlyOwner {
        IERC20(token_).transfer(wallet, amount);
    }

    function withDrawBNB(address wallet) external onlyOwner {
        payable(wallet).transfer(address(this).balance);
    }

    function checkUserOATCid(address addr, uint cid_) public view returns (uint[] memory){
        uint[] memory list;
        uint balance = OAT.balanceOf(addr);
        if (balance == 0) {
            return list;
        }
        uint amount;
        uint id;
        if (isOld[cid_]) {
            for (uint i = 0; i < balance; i ++) {
                id = OAT.tokenOfOwnerByIndex(addr, i);
                if (OAT.cid(id) == cid_) {
                    amount++;
                }
            }
            list = new uint[](amount);
            for (uint i = 0; i < balance; i ++) {
                id = OAT.tokenOfOwnerByIndex(addr, i);
                if (OAT.cid(id) == cid_) {
                    amount--;
                    list[amount] = id;
                }
            }
            return list;
        } else {
            for (uint i = 0; i < balance; i ++) {
                id = OAT.tokenOfOwnerByIndex(addr, i);
                if (info.cid(id) == cid_) {
                    amount++;
                }
            }
            list = new uint[](amount);
            for (uint i = 0; i < balance; i ++) {
                id = OAT.tokenOfOwnerByIndex(addr, i);
                if (info.cid(id) == cid_) {
                    amount--;
                    list[amount] = id;
                }
            }
            return list;
        }

    }

    function checkUserCid(address NFTAddr, address addr, uint cid_) public view returns (uint[] memory){
        uint[] memory list;
        uint balance = I721(NFTAddr).balanceOf(addr);
        if (balance == 0) {
            return list;
        }
        uint amount;
        uint id;
        for (uint i = 0; i < balance; i ++) {
            id = I721(NFTAddr).tokenOfOwnerByIndex(addr, i);
            if (I721(NFTAddr).cid(id) == cid_) {
                amount++;
            }
        }
        list = new uint[](amount);
        for (uint i = 0; i < balance; i ++) {
            id = I721(NFTAddr).tokenOfOwnerByIndex(addr, i);
            if (I721(NFTAddr).cid(id) == cid_) {
                amount--;
                list[amount] = id;
            }
        }
        return list;
    }


}