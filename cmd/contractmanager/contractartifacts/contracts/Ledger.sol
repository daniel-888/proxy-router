//SPDX-License-Identifier: MIT

pragma solidity >0.8.0;

contract Ledger {
    mapping(address => address) RentalContracts;
    address[] contractList;

    function setAddContractToStorage(address _rentalContract, address _seller) external {
        RentalContracts[_seller] = _rentalContract;
        contractList.push(_rentalContract);
    }

    function getListOfContractsLedger() external view returns (address[] memory) {
        return contractList;
    }
}