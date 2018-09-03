pragma solidity ^0.4.24;

contract StoreA {
    event ItemSet(uint256 key, address value);

    string public version;
    mapping (uint256 => address) public items;

    constructor(string _version) public {
        version = _version;
    }

    function setItem(uint256 key, address value) public {
        require(key >= 0);
        require(value != 0x0);
        items[key] = value;
        emit ItemSet(key, value);
    }

    function getItem(uint256 key) public view returns(address){
        require(key >= 0);
        return items[key];
    }
}
