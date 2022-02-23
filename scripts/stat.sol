pragma solidity >=0.5.0;

contract stat  {
    
    uint public txNum;
    uint64 public startTime;
    
    constructor(uint64 _startTime) public {
        startTime = _startTime;
    }

    function add() public {
        txNum += 1;
    }
}