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

