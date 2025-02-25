- switch 语句使用的表达式是单个字符，与之比较的 case 也是单个字符，因此不允许编译器使用 case “==” 的形式，即字节类型与 == 之类的字符串不能比较

- REPL (Read-Eval-Print Loop) 读取求值打印循环，很多语言都有REPL如：（Python，Ruby，JavaScript运行时）**REPL 读取输入将其发送到解释器进行求值，接着打印解释器的输出最后重新开始，重复这个循环**

- 词法分析器是一个软件组件，用来将输入的数据（通常是文本形式）解析成一个数据结构，通常是某种解析树，抽象语法树。将输入内容以结构化形式表示，在此过程中检查语法是否正确。

- 大多数解释器和编译器中，用于源代码内部表示的数据结构称为 “语法树” 或者 “抽象语法树”（AST：Abstract Syntax Tree）。

- CFG：上下文无关文法（context-free grammar）是一组规则，描述了如何根据一种语言的语法构成正确的语句。CFG最常用的符号格式是 Backus-Naur Form（BNF）或 Extended Backus-Naur Form（EBNF）。

- 变成语言进行语法分析时，主要有两种策略：自上而下的分析或自下而上的分析。每种策略都有很多变体。例如递归下降分析、Earley分析、预测分析、这些都是自上而下分析的变体。

- 自上而下的语法分析器和自下而上的语法分析器区别：前者从构造AST的根结点开始，然后下降；而后者则以相反的方式进行构造。

- 表达式会产生值，语句不会。

- 语法分析器部分有一个巧妙的小特性，即收集错误。语法分析器不会在遇到第一个错误时就退出，因此运行一次就可以捕捉所有语法错误，不用每次重新解析过程。

- 语法分析器的核心在于分析表达式，词法单元位置的有效性，此事普拉特解析法就派上用场了

- 普拉特解析法与其它语法分析方法主要区别是，普拉特没有将解析函数与语法规则（在BNF和EBNF中定义）相关联。这个想法的关键是，每种词法单元类型都可以具有两个与之相关联的解析函数，具体取决于词法单元的位置，比如是中缀还是前缀。
- 普拉特解析法是通过每个 token 的左右绑定力（lbp 和 rbp）来决定优先级

- 前缀运算符是位于操作数前面的运算符。like `--5`
- 后缀运算符是位于操作数后面的运算符。like `foobar++`
- 中缀运算符是未知两个操作数之间。like `1+2`

- 运算符优先级这个术语也可以称为运算顺序。它表示不同运算符的重要程度，能让运算符优先级更加直观。

- 将运算符优先级视为运算符黏性：即运算符黏住了周围多少个操作数。

- 普拉特语法分析器主要思想是将解析函数（普拉特称为语义代码）与词法单元类型相关联。每当遇到某个词法单元时，都会调用相关联的解析函数来解析对应的表达式，最后返回生成的 AST 节点。

- 前缀表达式在左侧是空的，中缀表达式左侧是有表达式的，需要记录

- 标识符无论是在耽搁表达式语句中，还是在其他上下文中，都是表达式：有些标识符是函数的调用参数；有些是中缀表达式操作数；还有些是条件中的单个表达式。

- 在区分运算符优先级的高低，使用了 iota。将不同常量设置逐个递增的数值。空白标识符_为0，其余的常量值是1到7.常量使用的数字无关紧要，重要的是顺序的彼此之间的关系。这些常量使用来区分运算符优先级的