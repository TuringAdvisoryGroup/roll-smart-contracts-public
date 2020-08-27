pragma solidity 0.5.0;

import "../util/Ownable.sol";
import "../util/SafeMath.sol";
import "../token/CreatorCoin.sol";
import "./TokenVesting.sol";

/**
 * @title TokenFactory contract for creating tokens from token proposals
 * @dev For creating tokens from pre-set parameters. This can be understood as a contract factory.
 * @author Jake Goh Si Yuan @ jakegsy, jake@jakegsy.com
 */
contract TokenFactory is Ownable{

    using SafeMath for uint256;

    uint8 public PLATFORM_PERCENTAGE;
    address public PLATFORM_WALLET; 
    TokenVesting public TokenVestingInstance;

    event LogTokenCreated(string name, string symbol, address indexed token, address vestingBeneficiary);
    event LogPlatformPercentageChanged(uint8 oldP, uint8 newP);
    event LogPlatformWalletChanged(address oldPW, address newPW);
    event LogTokenVestingChanged(address oldTV, address newTV);
    event LogTokenFactoryMigrated(address newTokenFactory);

    /**
     * @dev Constructor method
     * @param _tokenVesting address Address of tokenVesting contract. If set to address(0), it will create one instead.
     * @param _turingWallet address Turing Wallet address for sending out proportion of tokens alloted to it.
     */
    constructor(
        address _tokenVesting,
        address _turingWallet,
        uint8 _platformPercentage
    )
    validatePercentage(_platformPercentage)
    validateAddress(_turingWallet)
    public
    {
        
        require(_turingWallet != address(0), "Turing Wallet address must be non zero");
        PLATFORM_WALLET = _turingWallet;
        PLATFORM_PERCENTAGE = _platformPercentage;
        if (_tokenVesting == address(0)){
            TokenVestingInstance = new TokenVesting();
        }
        else{
            TokenVestingInstance = TokenVesting(_tokenVesting);
        }

    }


    /**
     * @dev Create token method
     * @param _name string Name parameter of Token
     * @param _symbol string Symbol parameter of Token
     * @param _decimals uint8 Decimals parameter of Token, restricted to < 18
     * @param _totalSupply uint256 Total Supply paramter of Token
     * @param _initialPercentage uint8 Initial percentage of total supply that the Vesting Beneficiary will receive from launch, restricted to < 100
     * @param _vestingPeriodInWeeks uint256 Number of weeks that the remaining of total supply will be linearly vested for, restricted to > 1
     * @param _vestingBeneficiary address Address of the Vesting Beneficiary 
     * @return address Address of token that has been created by those parameters
     */
    function createToken(
        string memory _name,
        string memory _symbol,
        uint8 _decimals,
        uint256 _totalSupply,
        uint8 _initialPercentage,
        uint256 _vestingPeriodInWeeks,
        address _vestingBeneficiary
        
    )
    public
    onlyOwner
    returns (address token)
    {
        uint256[3] memory proportions = calculateProportions(_totalSupply, _initialPercentage);
        require(proportions[0].add(proportions[1]).add(proportions[2]) == _totalSupply, 
        "The supply must be same as the proportion, sanity check.");
        CreatorCoin cc = new CreatorCoin(
            _name,
            _symbol,
            _decimals,
            proportions,
            _vestingBeneficiary,
            PLATFORM_WALLET,
            address(TokenVestingInstance)
        );
        TokenVestingInstance.addToken(address(cc), _vestingBeneficiary, _vestingPeriodInWeeks);
        token = address(cc);
        emit LogTokenCreated(_name, _symbol, token, _vestingBeneficiary);
    }

    /**
     * @dev Calculate proportions method
     * @param _totalSupply uint256 Total Supply paramter of Token
     * @param _initialPercentage uint8 Initial percentage of total supply that the Vesting Beneficiary will receive from launch, restricted to < 100
     * @dev Calculates supply given to the Turing platform, the Creator and the vesting supply
     * @return bytes32 Hash Index which is composed by the keccak256(name, symbol, msg.sender)
     */
    function calculateProportions(
        uint256 _totalSupply,
        uint8 _initialPercentage
    )
    private
    view
    validateTotalPercentage(_initialPercentage)
    returns (uint256[3] memory proportions)
    {
        proportions[0] = (_totalSupply).mul(_initialPercentage).div(100); //Initial Supply to Creator
        proportions[1] = (_totalSupply).mul(PLATFORM_PERCENTAGE).div(100); //Supply to Platform
        proportions[2] = (_totalSupply).sub(proportions[0]).sub(proportions[1]); // Remaining Supply to vest on
    }

    
    
    function setPlatformPercentage(
        uint8 _newPercentage
    )
    external
    validatePercentage(_newPercentage)
    onlyOwner
    {
        emit LogPlatformPercentageChanged(PLATFORM_PERCENTAGE, _newPercentage);
        PLATFORM_PERCENTAGE = _newPercentage;
    }
    
    function setPlatformWallet(
        address _newPlatformWallet
    )
    external
    validateAddress(_newPlatformWallet)
    onlyOwner
    {
        emit LogPlatformWalletChanged(PLATFORM_WALLET, _newPlatformWallet);
        PLATFORM_WALLET = _newPlatformWallet;
    }
    
    function migrateTokenFactory(
        address _newTokenFactory
    )
    external
    onlyOwner
    {
        TokenVestingInstance.transferOwnership(_newTokenFactory);
        emit LogTokenFactoryMigrated(_newTokenFactory);
    }

    function setTokenVesting(
        address _newTokenVesting
    )
    external
    onlyOwner
    {
        require(getOwnerStatic(_newTokenVesting) == address(this), "new TokenVesting not owned by TokenFactory");
        emit LogTokenVestingChanged(address(TokenVestingInstance), address(_newTokenVesting));
        TokenVestingInstance = TokenVesting(_newTokenVesting);
    }
    

    
    modifier validatePercentage(uint8 percentage){
        require(percentage > 0 && percentage < 100);
        _;
    }
    
    modifier validateAddress(address addr){
        require(addr != address(0));
        _;
    }

    modifier validateTotalPercentage(uint8 _x) {
        require(PLATFORM_PERCENTAGE + _x < 100);
        _;
    }

    function getTokenVesting() external view returns (address) {
        return address(TokenVestingInstance);
    } 
}


