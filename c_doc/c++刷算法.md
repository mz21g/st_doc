## C++特有的用const定义常量

C语言会用`#define`定义常量，但是C++里面用`const`这个限定符定义常量，这样做有个好处就是可以定义常量的类型

```cpp
const int a = 9999999;
```

## C++里面的string类

`string`只能用`cin`和`cout`处理，无法用`scanf`和`printf`处理

```c++
string s = "Hello world";
string s2 = s;
string s3 = s + s2;
string s4;
cin >> s4;
cout << s4;
```

用`cin`读入字符串时,是以空格为分隔符的,如果想读入一整行的字符串,就需要用`getline`

`s`的长度可以用`s.length()`获取, 有多少字符就是长度多少

```c++
string s;
getline(cin, s);
cout << s.length();
```

`string`中有个很常用的函数叫做`substr`,作用是截取某个字符串的子串

```c++
string s2 = s.substr(4); // 表示从下表4开始一直到结束
string s3 = s.substr(5, 3); // 表示从下标5开始, 3个字符
```

## C++ STL之动态数组vector的使用

它在头文件`vector`里面, 也在命名空间`std`里面, 所以使用的时候要引入头文件`#include <vector>`和`using namespace std;`

`vector`、`stack`、`queue`、`map`、`set`这些在C++中都叫做容器，这些容器的大小都可以用`.size()`获取到

`vector`可以一开始不定义大小，之后用`resize`方法分配大小，也可以一开始就定义大小，之后还可以对他插入删除动态改变它的大小

```c++
#include <iostream>
#include <vector>
using namespace std;
int main(){
  vector<int> v1; // 定义的时候没有分配大小
  cout << v1.size(); // 此处应该为0
  
  vector<int> v(10); // 直接定义长度为10的int数组，默认这10个元素值都为0
  // 或者
  vector<int> v2;
  v2.resize(8); // 先定义一个vector变量v2, 然后将长度resize为8，默认这8个元素值都是0
  
  // 在定义的时候就可以对vector变量进行初始化
  vector<int> v3(100, 9); // 把100长度的数组中所有的值都初始化为9
  
  // 访问的时候像数组一样直接使用[]下标访问即可
  v3[1] = 2;
  cout << v3[0];
  
  // 常用函数
  v3.push_bask(20); // 在vector v3的末尾添加一个元素20
  // 使用迭代器方式访问vector, c.end()指向容器的最后一个元素的后一个位置
  for (auto it = v3.begin(); it != c.end(); it++)
    cout << *it << " ";
  cout << endl;
  
  return 0;
}
```

## C++ STL之集合set的使用

`set`是集合，一个`set`里面的各个元素是各不相同的，而且`set`会按照元素进行从小到大排序

```c++
#include <iostream>
#include <set>
using namespace std;
int main(){
  set<int> s;
  s.insert(1); 
  cout << *(s.begin()) << endl; // 输出集合s的第一个元素（前面的星号表示要对指针取值）
  // 用迭代器遍历集合s里面的每一个元素
  for (auto it = s.begin(); it != s.end(); it++)
    cout << *it << " ";
  // 查找集合s中的值，如果结果等于s.end()表示未找到（因为s.end()表示s的最后一个元素的下一个元素所在的位置）
  cout << endl << (s.find(2) != s.end()) << endl;
  s.erase(2); // 删除集合s中的2这个元素
}
```

## C++ STL之映射map的使用

`map`使用的头文件`#include <map>`

`map`是键值对，比如一个人名对应一个学号，就可以定义一个字符串`string`类型的人名为”键“，学号`int`类型为”值“，如`map<string, int> m`，`map`会自动将所有的键值对从小到大排序

```c++
#include <iostream>
#include <map>
#include <string>
using namespace std;
int main(){
  map<string, int> m;
  m["hello"] = 2;
  cout << m["hello"] << endl; // 访问map的key为”hello“的value，如果key不存在，则返回0
  // 用迭代器遍历，输出map中所有的元素，键用it->first获取，值用it->second获取
  for (auto it = m.begin(); it != m.end(); it++)
    cout << it->first << " " << it->second << endl;
  // 访问map的第一个元素，输出它的键和值
  cout << m.begin()->first << " " << m.begin()->second << endl;
  // 访问map的最后一个元素，输出它的键和值
  cout << m.rbegin()->first << " " << m.rbegin()->second << endl;
  // 输出map的元素个数
  cout << m.size() << endl;
}
```

## C++ STL之unordered_map和unordered_set的使用

`unordered_map`在头文件`#include <unordered_map>`中，`unordered_set`在头文件`#include <unordered_set>`中

`map`会按照键值对的键`key`进行排序（`set`里面会按照集合中的元素大小进行排序，从小到大顺序）

`unordered_map`和`unordered_set`省去了这个排序的过程

## C++ STL之栈stack的使用

栈`stack`在头文件`#include <stack>`中

```c++
#include <iostream>
#include <stack>
using namespace std;
int main(){
  stack<int> s;
  for (int i = 0; i < 6; i++)
    s.push(i); // 压栈
  cout << s.top() << endl; // 访问栈顶元素
  cout << s.size() << endl; // 输出s的元素个数
  s.pop(); // 移除栈顶元素
  return 0;
}
```

## C++ STL之队列queue的使用

队列`queue`在头文件`#include <queue>`中

```c++
#include <iostream>
#include <queue>
using namespace std;
int main(){
  queue<int> q;
  for (int i = 0; i < 6; i++)
    q.push(i); // 压入队列
  cout << q.front() << " " << q.back() << endl; // 访问队列的队首元素和队尾元素
  cout << q.size() << endl; // 输出队列的元素个数
  q.pop(); // 移除队列的队首元素
  return 0;
}
```

