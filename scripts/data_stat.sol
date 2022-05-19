pragma solidity >=0.5.0;

contract dataStat {
    uint index;
    uint public txNum;
    uint64 public startTime;
    mapping(uint => bytes) data;

    function reset(uint64 _startTime) public {
        startTime = _startTime;
        txNum = 0;
    }

    function costManyGas(bytes memory input, uint64 complexity) public {
        for (uint i=0; i< complexity; i++) {
            index += 1;
            data[index] = input;
        }
        txNum += 1;
    }
}
