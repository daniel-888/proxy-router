//SPDX-License-Identifier: MIT

pragma solidity >0.8.0;

import "./Initializable.sol";

contract Implementation is Initializable{
    uint price;
    uint limit;
    uint speed;
    uint length;
    address seller;
    address validator;

    function initialize (
        uint _price,
        uint _limit,
        uint _speed,
        uint _length,
        address _seller,
        address _validator
    ) external initializer {
        price = _price;
        limit = _limit;
        speed = _speed;
        length = _length;
        seller = _seller;
        validator = _validator;
    }

    function getContractVariables() view external returns (uint[] memory) {
        uint[] memory hashrateParamenters;
        hashrateParamenters[0] = price;
        hashrateParamenters[1] = limit;
        hashrateParamenters[2] = speed;
        hashrateParamenters[3] = length;
        return hashrateParamenters;
    }
}