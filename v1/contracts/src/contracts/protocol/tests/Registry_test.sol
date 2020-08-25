pragma solidity 0.5.0;
pragma experimental ABIEncoderV2;

import "../Registry.sol";

contract RegistryTest {

    Registry public registry;
    
    event Success();

    constructor() public{
        registry = new Registry();
    }

    function testSubmitProposal() public returns (bool){
        bytes32 hashIndex = registry.submitProposal("JakeCoin1", "JC1", 11, 1000000, 15, 20, msg.sender, msg.sender);
        require(hashIndex == keccak256(abi.encodePacked("JakeCoin1", "JC1", msg.sender)), "Hash Index incorrect");


        Registry.Creator memory c = registry.getCreatorByIndex(hashIndex);
        require(c.token == address(0x0), "Address not 0");
        require(keccak256(abi.encodePacked(c.name)) == keccak256("JakeCoin1"), "Incorrect name");
        require(keccak256(abi.encodePacked(c.symbol)) == keccak256("JC1"), "Incorrect symbol");
        require(c.decimals == 11, "Incorrect decimals");
        require(c.totalSupply == 1000000, "Incorrect total supply");
        require(c.proposer == msg.sender, "Incorrect proposer");
        require(c.vestingBeneficiary == msg.sender, "Incorrect vesting beneficiary");
        require(c.initialPercentage == 15, "Incorrect initial percentrage");
        require(c.vestingPeriodInWeeks == 20, "Incorrect vesting period");
        require(c.approved == false, "Incorrect approved");

        emit Success();
        return true;
    }

    function testApproveProposal() public returns (bool){


        bytes32 hashIndex = registry.submitProposal("JakeCoin2", "JC2", 11, 1000000, 15, 20, msg.sender, msg.sender);
        Registry.Creator memory c = registry.getCreatorByIndex(hashIndex);
        require(c.token == address(0), "Incorrect token address before approval");
        require(c.approved == false, "Incorrect approved before approval");
        registry.approveProposal(hashIndex, address(69));
        c = registry.getCreatorByIndex(hashIndex);
        require(c.token == address(69), "Incorrect token address after approval");
        require(c.approved == true, "Incorrect approved after approval");

        emit Success();
        return true;
    }
}