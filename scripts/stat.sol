pragma solidity >=0.5.0;

contract stat  {
    
    uint public txNum;
    uint64 public startTime;
    
    constructor() public {
    }

	function reset(uint64 _startTime) public {
		startTime = _startTime;
		txNum = 0;
	}

    function add() public {
        txNum += 1;
    }
}