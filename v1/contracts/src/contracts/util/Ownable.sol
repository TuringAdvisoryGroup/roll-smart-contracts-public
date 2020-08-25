pragma solidity 0.5.0;


/**
 * @title Ownable
 * @dev The Ownable contract has an owner address, and provides basic authorization control
 * functions, this simplifies the implementation of "user permissions".
 */
contract Ownable {
    address public owner;


    event OwnershipTransferred(address indexed previousOwner, address indexed newOwner);


  /**
   * @dev The Ownable constructor sets the original `owner` of the contract to the sender
   * account.
   */
    constructor() public {
        owner = msg.sender;
    }

  /**
   * @dev Throws if called by any account other than the owner.
   */
    modifier onlyOwner() {
        require(msg.sender == owner);
        _;
    }

  /**
   * @dev Allows the current owner to transfer control of the contract to a newOwner.
   * @param newOwner The address to transfer ownership to.
   */
    function transferOwnership(address newOwner) public onlyOwner {
        require(newOwner != address(0));
        emit OwnershipTransferred(owner, newOwner);
        owner = newOwner;
    }
    
    function getOwner() public view returns (address) {
        return owner;
    }
    
    function getOwnerStatic(address ownableContract) internal view returns (address) {
        bytes memory callcodeOwner = abi.encodeWithSignature("getOwner()");
        (bool success, bytes memory returnData) = address(ownableContract).staticcall(callcodeOwner);
        require(success, "input address has to be a valid ownable contract");
        return parseAddr(returnData);
    }
    
    function getTokenVestingStatic(address tokenFactoryContract) internal view returns (address) {
        bytes memory callcodeTokenVesting = abi.encodeWithSignature("getTokenVesting()");
        (bool success, bytes memory returnData) = address(tokenFactoryContract).staticcall(callcodeTokenVesting);
        require(success, "input address has to be a valid TokenFactory contract");
        return parseAddr(returnData);
    }
    
    
    function parseAddr(bytes memory data) public pure returns (address parsed){
        assembly {parsed := mload(add(data, 32))}
    }
    
    


}