1. 在使用`gitline(cin, input);`时，如果前面有输入的话，要记得在`getline(cin, input);`前加一个`getchar();`

2. `s.insert(0, 4-s.length(), ‘0’);`用来给不足4位的时候补0

3. 四舍五入函数 `round()`

   ```cpp
   #include <iostream>
   #include <cmath>
   using namespace std;
   int main(){
     cout << round(1.5) << endl;
     // 保留一位小数，先乘以10.0，再除以10.0
     cout << round(10 / 8.0 * 10) / 10.0 << endl;
     return 0;
   }
   ```

4. `isalnum();` 用来判断一个字符是否为数字或字母