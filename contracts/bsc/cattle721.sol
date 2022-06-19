// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Enumerable.sol";
import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/utils/Strings.sol";
contract Star_Cattle is ERC721Enumerable, Ownable {
    using Strings for uint256;
    string public myBaseURI;
    uint public currentId;
    address public superMinter;
    uint male;
    uint female;
    uint limit;
    uint public maxLife;
    uint public burned;
    uint public maleCreation;
    uint public femaleCreation;
    uint[] public cowLimit;
    uint[] public starLimit;
    mapping(address => bool) public admin;
    mapping(address => uint) public mintersNormal;
    mapping(address => uint) public mintersCreation;
    mapping(address => mapping(address => mapping(uint => bool))) public isApprove;
    uint randomSeed;
    event BornNormalBull(address indexed sender_, uint indexed tokenId, uint life_, uint energy_, uint grow, uint star, uint[2] parents, uint deadTime, uint attack_, uint defense_, uint stamina_);

    event BornNormalCow(address indexed sender_, uint indexed tokenId, uint life_, uint energy_, uint grow, uint star, uint[2] parents, uint deadTime, uint milk_, uint milkRate_);

    event BornCreationBull(address indexed sender_, uint indexed tokenId, uint indexed creationId, uint life_, uint energy_, uint grow, uint deadTime, uint attack_, uint defense_, uint stamina_);

    event BornCreationCow(address indexed sender_, uint indexed tokenId, uint indexed creationId, uint life_, uint energy_, uint grow, uint deadTime, uint milk_, uint milkRate_);

    event AddDeadTime(uint indexed tokenId, uint indexed time);
    
    event GrowUp(uint indexed tokenId);

    event UpGradeStar(uint indexed tokenId,uint indexed stars);
    
     constructor() ERC721('Test Cattle', 'TCattle') {
         currentId = 1;
         superMinter = _msgSender();
         male = 10000;
         female = 10000;
         limit = 500;
         maxLife = 35 days;
         cowLimit = [5000,15000];
         starLimit = [25,40,60,80];
         randomSeed = 1;
         myBaseURI = 'https://ipfs.io/ipfs/bafybeicez4pkisfvvhkrbckcmip75tsi2ko5o2fy7y4dh62v5nz2iczyfe';
     }
    function rand(uint256 _length) internal returns (uint256) {
        uint256 random = uint256(keccak256(abi.encodePacked(block.difficulty, block.timestamp, randomSeed)));
        randomSeed ++;
        return random % _length + 1;
    }
    

    struct CowInfo {
        uint gender; // 1 for male,2 for femal;
        uint bornTime;
        uint energy;
        uint life;
        uint growth;
        uint exp;
        bool isAdult;
        uint attack;
        uint stamina;
        uint defense;
        uint milk;
        uint milkRate;
        uint star;
        uint[2] parents;
        uint deadTime;
    }


    mapping(uint => CowInfo) public cowInfoes;
    mapping(uint => bool) public isCreation;
    
    uint public creationAmount;
    mapping(uint => uint) public creationIndex;
    modifier onlyAdmin(){
        require(admin[_msgSender()], 'not admin');
        _;
    }
    function safeTransferFrom(
        address from,
        address to,
        uint256 tokenId
    ) public virtual override {
        //solhint-disable-next-line max-line-length
        require(_isApprovedOrOwner(_msgSender(), tokenId), "ERC721: transfer caller is not owner nor approved");
        if (!isCreation[tokenId]) {
            require(admin[msg.sender], 'not admin');
        }
        _transfer(from, to, tokenId);

    }

    function transferFrom(
        address from,
        address to,
        uint256 tokenId
    ) public virtual override {
        //solhint-disable-next-line max-line-length
        require(_isApprovedOrOwner(_msgSender(), tokenId), "ERC721: transfer caller is not owner nor approved");
        if (!isCreation[tokenId]) {
            require(admin[msg.sender], 'not admin');
        }
        safeTransferFrom(from, to, tokenId, "");
    }

    function setMintersNormal(address addr_, uint amount_) external onlyOwner {
        mintersNormal[addr_] = amount_;
    }
    
    function initCreation() external onlyOwner{
        uint bull;
        uint cow;
        for(uint i = 1; i < currentId; i ++){
            if(isCreation[i]){
                if(cowInfoes[i].gender == 1){
                    bull ++;
                    creationIndex[i] = bull;
                }else{
                    cow ++;
                    creationIndex[i] = cow;
                }
                
            }
        }
    }
    
    
    function setLimit(uint[] memory limit_) external onlyOwner{
        cowLimit = limit_;
    }

    function setMintersCreation(address addr_, uint amount_) external onlyOwner {
        mintersCreation[addr_] = amount_;
    }

    function setSuperMinter(address addr) external onlyOwner {
        superMinter = addr;
    }

    function setAdmin(address addr_, bool com_) external onlyOwner {
        admin[addr_] = com_;
    }

    function approve(address to, uint256 tokenId) public virtual override {
        isApprove[_msgSender()][to][tokenId] = true;
    }


    function addExp(uint tokenId_, uint amount_) external onlyAdmin {
        cowInfoes[tokenId_].exp += amount_;
    }

    function setGender(uint male_, uint female_) external onlyOwner {
        male = male_;
        female = female_;
    }

    function burn(uint tokenId_) public returns (bool){
        require(_isApprovedOrOwner(_msgSender(), tokenId_), "burner isn't owner");
        burned += 1;

        _burn(tokenId_);
        return true;
    }
    
    function mintNormallWithParents(address player) public {
        uint tokenId = currentId;
        uint[2] memory parents;
        if (_msgSender() != superMinter) {
            require(mintersNormal[_msgSender()] > 0, 'no mint amount');
            mintersNormal[_msgSender()] -= 1;
        }
        cowInfoes[currentId].gender = rand(2);
        cowInfoes[currentId].life = 7500 + rand(4000);
        cowInfoes[currentId].bornTime = block.timestamp;
        cowInfoes[currentId].energy = 7500 + rand(4000);
        cowInfoes[currentId].growth = 7500 + rand(4000);
        cowInfoes[currentId].isAdult = false;
        cowInfoes[currentId].deadTime = block.timestamp + (maxLife * cowInfoes[currentId].life / 10000);
        if (cowInfoes[currentId].gender == 1) {
            cowInfoes[currentId].attack = 7500 + rand(4000);
            cowInfoes[currentId].stamina = 7500 + rand(4000);
            cowInfoes[currentId].defense = 7500 + rand(4000);
            emit BornNormalBull(player, currentId, cowInfoes[currentId].life, cowInfoes[currentId].energy, cowInfoes[currentId].growth, 0, parents, cowInfoes[currentId].deadTime, cowInfoes[currentId].attack, cowInfoes[currentId].defense, cowInfoes[currentId].stamina);


        } else {
            cowInfoes[currentId].milk = 7500 + rand(4000);
            cowInfoes[currentId].milkRate = 8000 + rand(4000);
            emit BornNormalCow(player, currentId, cowInfoes[currentId].life, cowInfoes[currentId].energy, cowInfoes[currentId].growth, 0, parents, cowInfoes[currentId].deadTime, cowInfoes[currentId].milk, cowInfoes[currentId].milkRate);
        }
        currentId ++;
        _mint(player, tokenId);
    }

    function mintNormall(address player, uint[2] memory parents) public {
        uint tokenId = currentId;
        require(cowInfoes[parents[0]].gender == 1 && cowInfoes[parents[1]].gender == 2,'wrong parents');
        if (_msgSender() != superMinter) {
            require(mintersNormal[_msgSender()] > 0, 'no mint amount');
            mintersNormal[_msgSender()] -= 1;
        }
        CowInfo storage bull = cowInfoes[parents[0]];
        CowInfo storage cow = cowInfoes[parents[1]];
        cowInfoes[currentId].gender = rand(2);
        cowInfoes[currentId].life = coutintRand((bull.life + cow.life) / 2 );
        cowInfoes[currentId].bornTime = block.timestamp; 
        cowInfoes[currentId].energy = coutintRand((bull.energy + cow.energy) / 2 );
        cowInfoes[currentId].growth = coutintRand((bull.growth + cow.growth) / 2 );
        cowInfoes[currentId].parents = parents;
        cowInfoes[currentId].deadTime = block.timestamp + (maxLife * cowInfoes[currentId].life / 10000);
        if (cowInfoes[currentId].gender == 1) {
            cowInfoes[currentId].attack = coutintRand(bull.attack);
            cowInfoes[currentId].stamina = coutintRand(bull.stamina);
            cowInfoes[currentId].defense = coutintRand(bull.defense);
            emit BornNormalBull(player, currentId, cowInfoes[currentId].life, cowInfoes[currentId].energy, cowInfoes[currentId].growth, 0, parents, cowInfoes[currentId].deadTime, cowInfoes[currentId].attack, cowInfoes[currentId].defense, cowInfoes[currentId].stamina);


        } else {
            cowInfoes[currentId].milk = coutintRand(cow.milk);
            cowInfoes[currentId].milkRate = coutintRand(cow.milk);
            emit BornNormalCow(player, currentId, cowInfoes[currentId].life, cowInfoes[currentId].energy, cowInfoes[currentId].growth, 0, parents, cowInfoes[currentId].deadTime, cowInfoes[currentId].milk, cowInfoes[currentId].milkRate);
        }
        currentId ++;
        _mint(player, tokenId);

    }
    function growUp(uint tokenId) external{
        require(admin[msg.sender],'not admin');
        require(!isCreation[tokenId],'creation can not grow');
        require(!cowInfoes[tokenId].isAdult,'adult is true');
        cowInfoes[tokenId].isAdult = true;
        emit GrowUp(tokenId);
    }
    
    function upGradeStar(uint tokenId) external {
        require(admin[msg.sender], 'not admin');
        require(cowInfoes[tokenId].star < 3,"can't upgrade");
        cowInfoes[tokenId].star ++;
        uint deadTimes = block.timestamp + (maxLife * cowInfoes[tokenId].life / 10000);
        cowInfoes[tokenId].deadTime = deadTimes;
        emit AddDeadTime(tokenId,deadTimes);
        emit UpGradeStar(tokenId,cowInfoes[tokenId].star);

    }

    function mint(address player) public {
        _mint(player, currentId);
        if (_msgSender() != superMinter) {
            require(mintersCreation[_msgSender()] > 0, 'no mint amount');
            mintersCreation[_msgSender()] -= 1;
        }
        cowInfoes[currentId].gender = randGender();
        cowInfoes[currentId].life = 8000 + rand(4000);
        cowInfoes[currentId].bornTime = block.timestamp;
        cowInfoes[currentId].energy = 8000 + rand(4000);
        cowInfoes[currentId].growth = 8000 + rand(4000);
        cowInfoes[currentId].isAdult = true;
        cowInfoes[currentId].deadTime = block.timestamp + 360000 days;
        if (cowInfoes[currentId].gender == 1) {
            cowInfoes[currentId].attack = 8000 + rand(4000);
            cowInfoes[currentId].stamina = 8000 + rand(4000);
            cowInfoes[currentId].defense = 8000 + rand(4000);
            maleCreation++;
            creationIndex[currentId] = maleCreation;
            emit BornCreationBull(player, currentId, maleCreation, cowInfoes[currentId].life, cowInfoes[currentId].energy, cowInfoes[currentId].growth, cowInfoes[currentId].deadTime, cowInfoes[currentId].attack, cowInfoes[currentId].defense, cowInfoes[currentId].stamina);

        } else {
            cowInfoes[currentId].milk = 8000 + rand(4000);
            cowInfoes[currentId].milkRate = 8000 + rand(4000);
            femaleCreation ++;
            creationIndex[currentId] = femaleCreation;
            emit BornCreationCow(player, currentId, femaleCreation, cowInfoes[currentId].life, cowInfoes[currentId].energy, cowInfoes[currentId].growth, cowInfoes[currentId].deadTime, cowInfoes[currentId].milk, cowInfoes[currentId].milkRate);
        }
        isCreation[currentId] = true;
        creationAmount ++;
        currentId ++;
        

    }
    
    function checkUserCowList(address player) public view returns(uint[] memory){
        uint tempBalance = balanceOf(player);
        uint[] memory list = new uint[](tempBalance);
        uint token;
        for (uint i = 0; i < tempBalance; i++) {
            token = tokenOfOwnerByIndex(player, i);
            list[i] = token;
        }
        return list;

    }

    function checkUserCowListType(address player,bool creation_) public view returns (uint[] memory){
        uint tempBalance = balanceOf(player);
        uint token;
        uint count;
        for (uint i = 0; i < tempBalance; i++) {
            token = tokenOfOwnerByIndex(player, i);
            if(creation_ && isCreation[token]){
                count++;
            }
            if(!creation_ && !isCreation[token]){
                count++;
            }
        }
        uint[] memory list = new uint[](count);
        uint index;
        for (uint i = 0; i < tempBalance; i++) {
            token = tokenOfOwnerByIndex(player, i);
            if(creation_ && isCreation[token]){
                list[index] = token;
                index ++;
            }
            if(!creation_ && !isCreation[token]){
                list[index] = token;
                index ++;
            }
        }
        return list;

    }

    function randGender() internal returns (uint gen){
        uint out = rand(male + female);
        if (out > male) {
            female --;
            gen = 2;
        } else {
            male --;
            gen = 1;
        }
    }

    function addDeadTime(uint tokenId, uint time_) external {
        require(admin[msg.sender], 'not admin');
        require(cowInfoes[tokenId].bornTime > 0, 'nonexistent token');
        require(!isCreation[tokenId], 'can not add creation Cattle');
        cowInfoes[tokenId].deadTime += time_;
        emit AddDeadTime(tokenId, cowInfoes[tokenId].deadTime);
    }
    
    function coutintRand(uint com_) internal returns(uint){
        uint out = com_ * (900 + rand(200)) / 1000;
        if(out < cowLimit[0]){
            return cowLimit[0];
        }
        if (out > cowLimit[1]){
            return cowLimit[1];
        }
        return out;
    }


    function tokenURI(uint256 tokenId_) override public view returns (string memory){
        require(_exists(tokenId_), "nonexistent token");
        return string(abi.encodePacked(myBaseURI,tokenId_.toString()));
    }

    function setBaseURI(string memory uri) external onlyOwner{
        myBaseURI = uri;
    }

    function getGender(uint tokenId_) external view returns (uint){
        return cowInfoes[tokenId_].gender;
    }

    function getEnergy(uint tokenId_) external view returns (uint){
        return cowInfoes[tokenId_].energy;
    }

    function getAdult(uint tokenId_) external view returns (bool){
        return cowInfoes[tokenId_].isAdult;
    }

    function getLife(uint tokenId_) external view returns (uint){
        return cowInfoes[tokenId_].life;
    }

    function getBronTime(uint tokenId_) external view returns (uint){
        return cowInfoes[tokenId_].bornTime;
    }

    function getGrowth(uint tokenId_) external view returns (uint){
        return cowInfoes[tokenId_].growth;
    }

    function getMilk(uint tokenId_) external view returns (uint){
        if(isCreation[tokenId_]){
            return cowInfoes[tokenId_].milk;
        }
        uint star = cowInfoes[tokenId_].star;
        return cowInfoes[tokenId_].milk * starLimit[star] / 100;
    }

    function getStar(uint tokenId_) external view returns(uint) {
        return cowInfoes[tokenId_].star;
    }

    function getMilkRate(uint tokenId_) external view returns (uint){
        if(isCreation[tokenId_]){
            return cowInfoes[tokenId_].milkRate;
        }
        uint star = cowInfoes[tokenId_].star;
        return cowInfoes[tokenId_].milkRate * starLimit[star] / 100;
    }
    function getAttack(uint tokenId_) external view returns (uint){
        if(isCreation[tokenId_]){
            return cowInfoes[tokenId_].attack;
        }
        uint star = cowInfoes[tokenId_].star;
        return cowInfoes[tokenId_].attack * starLimit[star] / 100;
    }
    function getStamina(uint tokenId_) external view returns (uint){
        if(isCreation[tokenId_]){
            return cowInfoes[tokenId_].stamina;
        }
        uint star = cowInfoes[tokenId_].star;
        return cowInfoes[tokenId_].stamina * starLimit[star] / 100;
    }
    function getDefense(uint tokenId_) external view returns (uint){
        if(isCreation[tokenId_]){
            return cowInfoes[tokenId_].defense;
        }
        uint star = cowInfoes[tokenId_].star;
        return cowInfoes[tokenId_].defense * starLimit[star] / 100;
    }
    function deadTime(uint tokenId_) external view returns(uint){
        return cowInfoes[tokenId_].deadTime;
    }
    function getCowParents(uint tokenId_) external view returns(uint[2] memory){
        return cowInfoes[tokenId_].parents;
    }

}