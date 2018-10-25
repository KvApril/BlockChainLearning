pragma solidity ^0.4.0;

contract test {
    struct Person {
        string name;
        uint age;
        address addr;
    }
    mapping(address => Person) public records;

    function test(address _addr,string _name,uint _age){
        records[_addr] = Person(_name,_age,_addr);
    }

    function updateAge(address _addr,uint _age) public {
        Person storage p = records[_addr];
        if(p.addr == _addr){
            p.age = _age;
        }
    }
}
