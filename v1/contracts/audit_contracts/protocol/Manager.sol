pragma solidity 0.5.0;
pragma experimental ABIEncoderV2;

import "../util/Ownable.sol";
import "./Registry.sol";
import "./TokenFactory.sol";

/**
 * FOR THE AUDITOR
 * This contract was designed with the idea that it would be owned by
 * another multi-party governance-like contract such as a multi-sig
 * or a yet-to-be researched governance protocol to be placed on top of
 */


/** 
 * @title Manager contract for receiving proposals and creating tokens 
 * @dev For receiving token proposals and creating said tokens from such parameters.
 * @dev State is separated onto Registry contract
 * @dev To set up a working version of the entire platform, first create TokenFactory,
 * Registry and FanCoin, then transfer ownership to the Manager contract. See the truffle
 * test, especially manager_test.js to understand how this would be done offline.
 * @author Jake Goh Si Yuan @jakegsy, jake@jakegsy.com
 */
contract Manager is Ownable {

    using SafeMath for uint256;
    
    Registry public RegistryInstance;
    TokenFactory public TokenFactoryInstance;

    event LogTokenFactoryChanged(address oldTF, address newTF);
    event LogRegistryChanged(address oldR, address newR);
    event LogManagerMigrated(address indexed newManager);
    
    /** 
     * @dev Constructor on Manager 
     * @param _registry address Address of Registry contract
     * @param _tokenFactory address Address of TokenFactory contract
     * @notice It is recommended that all the component contracts be launched before Manager
     */
    constructor(
        address _registry,
        address _tokenFactory
    )
    public
    {
        require(_registry != address(0) && _tokenFactory != address(0));
        TokenFactoryInstance = TokenFactory(_tokenFactory);
        RegistryInstance = Registry(_registry);
    }

    /**
     * @dev Submit Token Proposal
     * @param _name string Name parameter of Token
     * @param _symbol string Symbol parameter of Token
     * @param _decimals uint8 Decimals parameter of Token, restricted to < 18
     * @param _totalSupply uint256 Total Supply paramter of Token
     * @param _initialPercentage uint8 Initial percentage of total supply that the Vesting Beneficiary will receive from launch, restricted to < 100
     * @param _vestingPeriodInWeeks uint256 Number of weeks that the remaining of total supply will be linearly vested for, restricted to > 1
     * @param _vestingBeneficiary address Address of the Vesting Beneficiary 
     * @return bytes32 Hash Index which is composed by the keccak256(name, symbol, msg.sender)
     */

    function submitProposal(
        string memory _name,
        string memory _symbol,
        uint8 _decimals,
        uint256 _totalSupply,
        uint8 _initialPercentage,
        uint256 _vestingPeriodInWeeks,
        address _vestingBeneficiary
    )
    validatePercentage(_initialPercentage)
    validateDecimals(_decimals)
    validateVestingPeriod(_vestingPeriodInWeeks)
    isInitialized()
    public
    returns (bytes32 hashIndex)
    {
        hashIndex = RegistryInstance.submitProposal(_name,_symbol,_decimals,_totalSupply,
        _initialPercentage, _vestingPeriodInWeeks, _vestingBeneficiary, msg.sender);
    }
    
    /**
     * @dev Approve Token Proposal 
     * @param _hashIndex bytes32 Hash Index of Token Proposal, given by keccak256(name, symbol, msg.sender)
     */
    function approveProposal(
        bytes32 _hashIndex
    )
    isInitialized()
    onlyOwner
    external
    {   
        //Registry.Creator memory approvedProposal = RegistryInstance.rolodex(_hashIndex);
        Registry.Creator memory approvedProposal = RegistryInstance.getCreatorByIndex(_hashIndex);
        address ac = TokenFactoryInstance.createToken(
            approvedProposal.name,
            approvedProposal.symbol,
            approvedProposal.decimals,
            approvedProposal.totalSupply,
            approvedProposal.initialPercentage,
            approvedProposal.vestingPeriodInWeeks,
            approvedProposal.vestingBeneficiary
            );
        bool success = RegistryInstance.approveProposal(_hashIndex, ac);
        require(success, "Registry approve proposal has to succeed");
    }


    /*
     * CHANGE PLATFORM VARIABLES AND INSTANCES
     */
    

    function setPlatformWallet(
        address _newPlatformWallet
    )
    onlyOwner
    isInitialized()
    external
    {
        TokenFactoryInstance.setPlatformWallet(_newPlatformWallet);
    }

    function setTokenFactoryPercentage(
        uint8 _newPercentage
    )
    onlyOwner
    validatePercentage(_newPercentage)
    isInitialized()
    external
    {
        TokenFactoryInstance.setPlatformPercentage(_newPercentage);
    }
    
    function setTokenFactory(
        address _newTokenFactory
    )
    onlyOwner
    external
    {
        
        require(getOwnerStatic(_newTokenFactory) == address(this), "new TokenFactory has to be owned by Manager");
        require(getTokenVestingStatic(_newTokenFactory) == address(TokenFactoryInstance.TokenVestingInstance()), "TokenVesting has to be the same");
        TokenFactoryInstance.migrateTokenFactory(_newTokenFactory);
        require(getOwnerStatic(getTokenVestingStatic(_newTokenFactory))== address(_newTokenFactory), "TokenFactory does not own TokenVesting");
        emit LogTokenFactoryChanged(address(TokenFactoryInstance), address(_newTokenFactory));
        TokenFactoryInstance = TokenFactory(_newTokenFactory); 
    }
    
    function setRegistry(
        address _newRegistry
    )

    onlyOwner
    external
    {
        require(getOwnerStatic(_newRegistry) == address(this), "new Registry has to be owned by Manager");
        emit LogRegistryChanged(address(RegistryInstance), _newRegistry);
        RegistryInstance = Registry(_newRegistry);
    }
    
    function setTokenVesting(
        address _newTokenVesting
    )
    onlyOwner
    external
    {
        TokenFactoryInstance.setTokenVesting(_newTokenVesting);
    }

    function migrateManager(
        address _newManager
    )
    onlyOwner
    isInitialized()
    external
    {
        RegistryInstance.transferOwnership(_newManager);
        TokenFactoryInstance.transferOwnership(_newManager);
        emit LogManagerMigrated(_newManager);
    }

    modifier validatePercentage(uint8 percentage) {
        require(percentage > 0 && percentage < 100, "has to be above 0 and below 100");
        _;
    }

    modifier validateDecimals(uint8 decimals) {
        require(decimals > 0 && decimals < 18, "has to be above 0 and below 18");
        _;
    }

    modifier validateVestingPeriod(uint256 vestingPeriod) {
        require(vestingPeriod > 1, "has to be above 1");
        _;
    }

    modifier isInitialized() {
        require(initialized(), "manager not initialized");
        _;
    }

    function initialized() public view returns (bool){
        return (RegistryInstance.owner() == address(this)) &&
            (TokenFactoryInstance.owner() == address(this)) &&
            (TokenFactoryInstance.TokenVestingInstance().owner() == address(TokenFactoryInstance)); 
    }

    
    
}
