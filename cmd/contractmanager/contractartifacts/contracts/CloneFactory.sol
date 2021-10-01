//SPDX-License-Identifier: MIT

pragma solidity >0.8.0;

import "./Clones.sol";
import "./Ownable.sol";
import './Implementation.sol';

/// @title CloneFactory
/// @author Josh Kean (Lumerin)
/// @notice Variables passed into contract initializer are subject to change based on the design of the hashrate contract

contract CloneFactory {
  address baseImplementation; 
  address contractManager;
  address lmn;
  address proxy;

  constructor() {
    Implementation _imp = new Implementation();
    baseImplementation = address(_imp);
  }

  //function to create a new Implementation contract
  function setCreateNewRentalContract(
    uint _price,
    uint _limit,
    uint _speed,
    uint _length,
    address _seller
  ) external returns (address) {
    address _newContract = Clones.clone(baseImplementation); 
    Implementation(_newContract).initialize(_price, _limit, _speed, _length, _seller, contractManager, lmn);
    return _newContract;
  }


  function setChangeContractManagerAddress(address _contractManager) public { //add modifier to specify owner
    contractManager = _contractManager;
  }


  function setChangeProxyAddress(address _proxy) public { //add modifier to specify owner
    proxy = _proxy;
  }
}