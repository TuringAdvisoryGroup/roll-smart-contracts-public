pragma solidity 0.5.0;
pragma experimental ABIEncoderV2;

import "../Manager.sol";
import "../Registry.sol";
import "../TokenFactory.sol";
import "../TokenVesting.sol";
import "../../token/CreatorCoin.sol";


contract ManagerTest {

    Manager public ManagerInstance;
    Registry public RegistryInstance;
    TokenVesting public TokenVestingInstance;
    TokenFactory public TokenFactoryInstance;
    
    event Success();
    
    constructor() public{
        
        RegistryInstance = new Registry();
        TokenVestingInstance = new TokenVesting();
        TokenFactoryInstance = new TokenFactory(address(TokenVestingInstance), address(this), 5);
        ManagerInstance = new Manager(
            address(RegistryInstance), 
            address(TokenFactoryInstance));
            
        RegistryInstance.transferOwnership(address(ManagerInstance));
        TokenVestingInstance.transferOwnership(address(TokenFactoryInstance));
        TokenFactoryInstance.transferOwnership(address(ManagerInstance));
    }
    
    function testSetPlatformWallet(address newPlatform) public returns (bool){
        ManagerInstance.setPlatformWallet(newPlatform);
        require(TokenFactoryInstance.PLATFORM_WALLET() == newPlatform);
        emit Success();
        return true;
    }
    
    function testSetTokenFactoryPercentage(uint8 newPercentage) public returns (bool){
        ManagerInstance.setTokenFactoryPercentage(newPercentage);
        require(TokenFactoryInstance.PLATFORM_PERCENTAGE() == newPercentage);
        emit Success();
        return true;
    }
    
    // function testSetTokenFactory() public returns(bool){
    //     require(address(ManagerInstance.TokenFactoryInstance()) == address(TokenFactoryInstance));
    //     TokenFactory newTF = new TokenFactory(address(0), address(this), 5);
    //     ManagerInstance.setTokenFactory(newTF);
    //     require(address(ManagerInstance.TokenFactoryInstance()) == address(newTF));
    //     emit Success();
    //     return true;
    // }
    
    // function testSetRegistry() public returns(bool){
    //     require(address(ManagerInstance.RegistryInstance()) == address(RegistryInstance));
    //     Registry newR = new Registry();
    //     ManagerInstance.setRegistry(newR);
    //     require(address(ManagerInstance.RegistryInstance()) == address(newR));
    //     emit Success();
    //     return true;
    // }
    
    function testMigrateManager() public returns(bool){
        Manager newM = new Manager(address(RegistryInstance), address(TokenFactoryInstance));
        require(!newM.initialized());
        ManagerInstance.migrateManager(address(newM));
        require(newM.initialized());
        require(RegistryInstance.owner() == address(newM));
        require(TokenFactoryInstance.owner() == address(newM));
        emit Success();
        ManagerInstance = newM;
        return true;
    }

    function testSubmitAndApproveProposal() public returns (bool){
        
        string memory name = "JakeCoin";
        string memory symbol = "JC";
        uint8 decimals = 10;
        uint256 totalSupply = 1*10**7;
        uint8 initialPercentage = 18;
        uint256 vestingPeriodInWeeks = 12;
        address vestingBeneficiary = msg.sender;
        
        
        bytes32 hashIndex = ManagerInstance.submitProposal(name, symbol, decimals, 
        totalSupply, initialPercentage, vestingPeriodInWeeks, vestingBeneficiary);
        
        require(hashIndex == keccak256(abi.encodePacked(name, symbol, address(this))) , "Hash Index incorrect");
        
        Registry.Creator memory c = RegistryInstance.getCreatorByIndex(hashIndex);
        
        //Checking whether state matches parameters
        require(c.token == address(0x0), "Address not 0");
        require(keccak256(abi.encodePacked(c.name)) == keccak256(abi.encodePacked(name)), 
        "Incorrect name at submit");
        require(keccak256(abi.encodePacked(c.symbol)) == keccak256(abi.encodePacked(symbol)), 
        "Incorrect symbol at submit");
        require(c.decimals == decimals, "Incorrect decimals at submit");
        require(c.totalSupply == totalSupply, "Incorrect total supply at submit");
        require(c.proposer == address(this), "Incorrect proposer at submit");
        require(c.vestingBeneficiary == vestingBeneficiary, "Incorrect vesting beneficiary at submit");
        require(c.initialPercentage == initialPercentage, "Incorrect initial percentrage at submit");
        require(c.vestingPeriodInWeeks == vestingPeriodInWeeks, "Incorrect vesting period at submit");
        require(c.approved == false, "Incorrect approved at submit");
        
        
        //Now let us start the approve part of the process
        ManagerInstance.approveProposal(hashIndex);
        c = RegistryInstance.getCreatorByIndex(hashIndex);
        require(c.approved == true, "Incorrect approved after approval");
        require(c.token != address(0), "Incorrect artist token address after approval");
        
        //Checking details of actual token
        CreatorCoin testCoin = CreatorCoin(c.token);
        require(keccak256(abi.encodePacked(testCoin.name())) == keccak256(abi.encodePacked(name)), "Incorrect token name");
        require(keccak256(abi.encodePacked(testCoin.symbol())) == keccak256(abi.encodePacked(symbol)), "Incorrect token symbol");
        require(testCoin.decimals() == decimals, "Incorrect token decimals");
        require(testCoin.totalSupply() == totalSupply, "Incorrect token totalSupply");
        
        
        TokenVesting.VestingInfo memory v_info = TokenVestingInstance.getVestingInfo(c.token);
        uint256 expectedVestingPeriod = c.vestingPeriodInWeeks * TokenVestingInstance.WEEKS_IN_SECONDS();
        require(v_info.vestingBeneficiary == tx.origin, "Incorrect vesting beneficiary");
        require(v_info.duration == expectedVestingPeriod, "Incorrect vesting period");
        emit Success();
        
        return true;
    }

        
        

    

    function changeOwnershipManager() public {
        ManagerInstance.transferOwnership(msg.sender);
    }

    
}