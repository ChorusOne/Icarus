pragma solidity ^0.5.3;


library SafeMath {
    
    function mul(uint256 a, uint256 b) internal pure returns (uint256) {
        
        
        
        if (a == 0) {
            return 0;
        }

        uint256 c = a * b;
        require(c / a == b);

        return c;
    }

    
    function div(uint256 a, uint256 b) internal pure returns (uint256) {
        
        require(b > 0);
        uint256 c = a / b;
        

        return c;
    }

    
    function sub(uint256 a, uint256 b) internal pure returns (uint256) {
        require(b <= a);
        uint256 c = a - b;

        return c;
    }

    
    function add(uint256 a, uint256 b) internal pure returns (uint256) {
        uint256 c = a + b;
        require(c >= a);

        return c;
    }

    
    function mod(uint256 a, uint256 b) internal pure returns (uint256) {
        require(b != 0);
        return a % b;
    }
}

contract Initializable {
  bool public initialized;

  modifier initializer() {
    require(!initialized);
    initialized = true;
    _;
  }
}

interface IERC20Token {
  function transfer(address, uint256) external returns (bool);
  function approve(address, uint256) external returns (bool);
  function transferFrom(address, address, uint256) external returns (bool);
  function totalSupply() external view returns (uint256);
  function balanceOf(address) external view returns (uint256);
  function allowance(address, address) external view returns (uint256);

  
  event Transfer(address indexed from, address indexed to, uint256 value);
  event Approval(address indexed owner, address indexed spender, uint256 value);
}

interface ICeloToken {
  function transferWithComment(address, uint256, string calldata) external returns (bool);
  function name() external view returns (string memory);
  function symbol() external view returns (string memory);
  function decimals() external view returns (uint8);
}

contract GoldToken is Initializable, IERC20Token, ICeloToken {
  using SafeMath for uint256;

  
  
  address constant TRANSFER = address(0xfd);
  string constant NAME = "Celo Gold";
  string constant SYMBOL = "cGLD";
  uint8 constant DECIMALS = 18;
  uint256 internal totalSupply_;
  

  mapping(address => mapping(address => uint256)) internal allowed;

  event Transfer(address indexed from, address indexed to, uint256 value);

  event TransferComment(string comment);

  event Approval(address indexed owner, address indexed spender, uint256 value);

  
  modifier onlyVm() {
    require(msg.sender == address(0), "sender was not vm (reserved 0x0 addr)");
    _;
  }

  
  
  function initialize() external initializer {
    totalSupply_ = 0;
  }

  
  
  function transfer(address to, uint256 value) external returns (bool) {
    return _transfer(to, value);
  }

  
  function transferWithComment(address to, uint256 value, string calldata comment)
    external
    returns (bool)
  {
    bool succeeded = _transfer(to, value);
    emit TransferComment(comment);
    return succeeded;
  }

  
  function approve(address spender, uint256 value) external returns (bool) {
    allowed[msg.sender][spender] = value;
    emit Approval(msg.sender, spender, value);
    return true;
  }

  
  function increaseAllowance(address spender, uint256 value) external returns (bool) {
    uint256 oldValue = allowed[msg.sender][spender];
    uint256 newValue = oldValue.add(value);
    allowed[msg.sender][spender] = newValue;
    emit Approval(msg.sender, spender, newValue);
    return true;
  }

  
  function decreaseAllowance(address spender, uint256 value) external returns (bool) {
    uint256 oldValue = allowed[msg.sender][spender];
    uint256 newValue = oldValue.sub(value);
    allowed[msg.sender][spender] = newValue;
    emit Approval(msg.sender, spender, newValue);
    return true;
  }

  
  function transferFrom(address from, address to, uint256 value) external returns (bool) {
    require(to != address(0), "transfer attempted to reserved address 0x0");
    require(value <= balanceOf(from), "transfer value exceeded balance of sender");
    require(
      value <= allowed[from][msg.sender],
      "transfer value exceeded sender's allowance for recipient"
    );

    bool success;
    (success, ) = TRANSFER.call.value(0).gas(gasleft())(abi.encode(from, to, value));
    require(success, "Celo Gold transfer failed");

    allowed[from][msg.sender] = allowed[from][msg.sender].sub(value);
    emit Transfer(from, to, value);
    return true;
  }

  
  function name() external view returns (string memory) {
    return NAME;
  }

  
  function symbol() external view returns (string memory) {
    return SYMBOL;
  }

  
  function decimals() external view returns (uint8) {
    return DECIMALS;
  }

  
  function totalSupply() external view returns (uint256) {
    return totalSupply_;
  }

  
  function allowance(address owner, address spender) external view returns (uint256) {
    return allowed[owner][spender];
  }

  
  function increaseSupply(uint256 amount) external onlyVm {
    totalSupply_ = totalSupply_.add(amount);
  }

  
  function balanceOf(address owner) public view returns (uint256) {
    return owner.balance;
  }

  
  function _transfer(address to, uint256 value) internal returns (bool) {
    require(to != address(0), "transfer attempted to reserved address 0x0");
    require(value <= balanceOf(msg.sender), "transfer value exceeded balance of sender");

    bool success;
    (success, ) = TRANSFER.call.value(0).gas(gasleft())(abi.encode(msg.sender, to, value));
    require(success, "Celo Gold transfer failed");
    emit Transfer(msg.sender, to, value);
    return true;
  }
}