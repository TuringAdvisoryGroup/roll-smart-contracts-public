pragma solidity 0.5.0;
pragma experimental ABIEncoderV2;


import "../util/SafeMath.sol";
import "../token/ERC20/ERC20.sol";
import "../util/Ownable.sol";

/**
 * @title TokenVesting contract for linearly vesting tokens to the respective vesting beneficiary
 * @dev This contract receives accepted proposals from the Manager contract, and holds in lieu
 * @dev all the tokens to be vested by the vesting beneficiary. It releases these tokens when called
 * @dev upon in a continuous-like linear fashion.
 * @notice This contract was written with reference to the TokenVesting contract from openZeppelin
 * @notice @ https://github.com/OpenZeppelin/openzeppelin-solidity/blob/master/contracts/drafts/TokenVesting.sol
 * @author Jake Goh Si Yuan @ jakegsy, jake@jakegsy.com
 */
contract TokenVesting is Ownable{

    using SafeMath for uint256;

    event Released(address indexed token, address vestingBeneficiary, uint256 amount);
    event LogTokenAdded(address indexed token, address vestingBeneficiary, uint256 vestingPeriodInWeeks);

    uint256 constant public WEEKS_IN_SECONDS = 1 * 7 * 24 * 60 * 60;

    struct VestingInfo {
        address vestingBeneficiary;
        uint256 releasedSupply;
        uint256 start;
        uint256 duration;
    }

    mapping(address => VestingInfo) public vestingInfo;

    /**
     * @dev Method to add a token into TokenVesting
     * @param _token address Address of token 
     * @param _vestingBeneficiary address Address of vesting beneficiary 
     * @param _vestingPeriodInWeeks uint256 Period of vesting, in units of Weeks, to be converted
     * @notice This emits an Event LogTokenAdded which is indexed by the token address
     */
    function addToken
    (
        address _token,
        address _vestingBeneficiary,
        uint256 _vestingPeriodInWeeks
    )
    external
    onlyOwner
    {
        vestingInfo[_token] = VestingInfo({
            vestingBeneficiary : _vestingBeneficiary,
            releasedSupply : 0,
            start : now,
            duration : uint256(_vestingPeriodInWeeks).mul(WEEKS_IN_SECONDS)
        });
        emit LogTokenAdded(_token, _vestingBeneficiary, _vestingPeriodInWeeks);
    }

    /**
     * @dev Method to release any already vested but not yet received tokens
     * @param _token address Address of Token
     * @notice This emits an Event LogTokenAdded which is indexed by the token address
     */

    function release
    (
        address _token
    )
    external
    {   
        uint256 unreleased = releaseableAmount(_token);
        require(unreleased > 0);
        vestingInfo[_token].releasedSupply = vestingInfo[_token].releasedSupply.add(unreleased);
        bool success = ERC20(_token).transfer(vestingInfo[_token].vestingBeneficiary, unreleased);
        require(success, "transfer from vesting to beneficiary has to succeed");
        emit Released(_token, vestingInfo[_token].vestingBeneficiary, unreleased);
    }

    /**
     * @dev Method to check the quantity of token that is already vested but not yet received
     * @param _token address Address of Token
     * @return uint256 Quantity of token that is already vested but not yet received
     */
    function releaseableAmount
    (
        address _token
    )
    public
    view
    returns(uint256)
    {
        return vestedAmount(_token).sub(vestingInfo[_token].releasedSupply);
    }

    /**
     * @dev Method to check the quantity of token vested at current block
     * @param _token address Address of Token
     * @return uint256 Quantity of token that is vested at current block
     */

    function vestedAmount
    (
        address _token
    )
    public
    view
    returns(uint256)
    {
        VestingInfo memory info = vestingInfo[_token];
        uint256 currentBalance = ERC20(_token).balanceOf(address(this));
        uint256 totalBalance = currentBalance.add(info.releasedSupply);
        if (now >= info.start.add(info.duration)) {
            return totalBalance;
        } else {
            return totalBalance.mul(now.sub(info.start)).div(info.duration);
        }

    }


    function getVestingInfo
    (
        address _token
    )
    external
    view
    returns(VestingInfo memory)
    {
        return vestingInfo[_token];
    }
    

}
