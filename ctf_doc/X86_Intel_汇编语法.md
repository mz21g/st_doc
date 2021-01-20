## x86寄存器体系

![image-20201208155328061](C:\Users\CLAY\Desktop\st_doc\ctf_doc\image\pwn\X86寄存器体系.png)

## 汇编基础

```assembly
; EAX EBX ECX EDX - 通用寄存器，虽各有用途，但主要还是用作CPU执行机器指令时的临时变量
; EAX - 累加寄存器
; EBX - 计数寄存器
; ECX - 数据寄存器
; EDX - 基址寄存器
; 指针寄存器
; ESP - 堆栈指针，指向当前栈帧的栈顶（底地址）
; EBP - 基指针，指向当前栈帧的栈底（高地址）
; ESI 和 EDI 也是指针，当需要读取或写入数据时，通常使用他们指向源和目的地址。大多数情况下可视为通用寄存器
; ESI - 源变址
; EDI - 目的变址
; EIP - 指令指针，指向CPU要执行的下一条指令的地址
; 标志寄存器
; ZF - 零标志，当一个操作的结果为0时，设置为1
; CF - 进位标志，当一个操作的结果太大或者太小进而产生进位或借位时，设置为1
; SF - 负数标志，当一个操作的结果为负数时，设置为1
; DF - 方向标志位，表示DI以及SI自增（减）的偏移地址寄存器的自增（减）方向

; 移动数据
mov ebx, eax ; 将eax的值移动到ebx
mov eax, 0xDEADBEEF ; 将0xDEADBEEF移动到eax
mov edx, DWORD PTR [0X41424344] ; 将0X41424344地址的4字节值移动到edx
mov ecx, DWORD PTR [edx] ; 将地址为edx值位置的4字节值移动到ecx
mov eax, DWORD PTR [ecx+esi*8] ; 将地址为 ecx+esi*8值位置的4字节值移动到eax
lea eax, [ebx+b] ; 将ebx+b这个值直接赋给eax，而不是ebx+8处的内存地址里的数据赋给eax

; 算术运算
sub edx, 0x11 ; edx = edx - 0x11
add eax, ebx ; eax = eax + ebx
inc edx ; edx++
dec ebx ; ebx--
xor eax, eax ; eax = eax ^ eax
or edx, 0x1337 ; edx = edx | 0x1337

repne scasb ; 字符串扫描
;----------------------------------
; 实例:
section .data

	EditBuff: db 'abcdefghijklm#',10
	BUFFERLEN equ $-EditBuff
	FILLCHR	  equ 35 ;'#'
section .text
_start:
	cld                 ; 控制edi指向的字符串地址变化方向是从低到高
	mov al, FILLCHR     ; '#' -> al
	mov edi,EditBuff	; "abcdefghijklm" -> edi
	mov ecx,0000ffffh	; ecx = 65535 
	repne scasb         ; 扫描edi指向的字符串，扫描方向从低位地址向高位地址，如果遇到字节等于al或者ecx计数为0，则结束扫描。
	mov byte [edi-1], '$' ; 因为每次循环都会把edi加1，因此edi-1才是等于al那个字节（对于本例ecx肯定不会为0）。所以本条指令会把’#’替换成’$’
;----------------------------------

; 条件跳转
jz $LOC ; 如果 ZF=1, 跳转到 $LOC
jnz $LOC ; 如果 ZF=0, 跳转到 $LOC
jg $LOC ; 如果一个比较操作的结果，destination > source， 跳转到 $LOC

; 栈操作
push ebx ; 栈指针减去4，移动到一个较低的地址，并将ebx的值复制到栈的顶部
;----------------------------------
; push ebx 可以写成:
;	sub esp, 4
;	mov DWORD PTR [esp], ebx
;----------------------------------

pop ebx ; 从栈的顶部复制值给ebx，然后栈指针加4，将它移动到一个较高的地址
;----------------------------------
; pop ebx 可以写成:
;	mov ebx, DWORD PTR [esp]
;	add esp, 4
;----------------------------------


; 调用/返回（Calling/Returning）
call some_function 
;----------------------------------
; 调用some_function处的代码，我们需要将返回地址入栈
; 	push eip -- not actually valid
;----------------------------------

ret
;---------------------------------------------
; 从函数调用处返回，将栈顶的eip出栈
; 	pop eip -- not actually valid
;---------------------------------------------

nop ; 'no operation' - does nothing
```

## x86实战

```assembly
;---------------------------------------------
; BYTE: 字节，8位，用来储存char或者char类型指针
; WORD: 字, 16位，用来储存16为整数或者16位地址
; DWORD: 双字，32位，用来存储32位整数或32为地址
;---------------------------------------------
0x08048624: "HELLO WORLD\0"   ; 9个字节的字符串
	mov ebx, 0x08048624       ; char *ebx = "HELLO WORLD\0" 
	mov eax, 0                ; eax = 0
LOOPY:
	mov c1, BYTE PTR [ebx]    ; char c1 = *ebx
	cmp c1, 0                 ; c1 == 0
	jz end                    ; if c1 == 0, go to end
	inc eax                   ; eax++
	inc ebx                   ; ebx++
	jmp LOOPY                 ; goto LOOPY
end:
	ret                       ; return (eax位置的字符串的长度)
```

* **反汇编**

```c
char * word = "HELLO WORLD\0";
int len = 0;

while(*word != 0){
    len++;
    word++;
}

return len;
```

### IDA 转换后的机器代码详解

![97b1e7c9b464ac40b0ce4176da7a373](C:\Users\CLAY\Desktop\ida\97b1e7c9b464ac40b0ce4176da7a373.png)

这是一个小程序的main()函数经过ida转换后的机器代码，上图用红框框起来的部分每两个十六进制数就是一个字节，这里可能有人不太明白，我详细说一下。

一个字节包含8位，每一位可以是0或1，所以一个字节可以表示从 $(00000000)_2$ - $(11111111)_2$，共计256个可能值。又因为四位二进制可以表示一位的十六进制，比如$ (0000)_2 $ - 0x0，$ (1010)_2 $ - 0xA，$ (1100)_2 $ - 0xC，$ (1111)_2 $ - 0xF，所以$(00000000)_2$ 表示为 0x00，$(11111111)_2$表示为0xFF，所以一个字节又可以表示为从 0x00 - 0xFF，共计256个可能值。

红框框起来的左边部分，也就是以0x080484B4开头的十六进制数是内存地址，内存只是用地址进行编码的临时存储空间的字节的集合。我们可以将内存就看作是一排字节，每个字节都有自己的内存地址（一个内存地址只存放一个字节），可通过内存地址来访问内存中的每个字节。在ida中，只会显示一排中第一个字节的内存地址，比如第一排的0x55（内存地址为0x080484B4）、第二排的0x89（内存地址为0x080484B5），而第二排的0xE5（内存地址为0x080484B6）则不会显示出来。

汇编语言指令与其对应的机器指令存在直接的一一对应关系，这意味着，由于每个处理器体系结构都有不同的机器语言指令，所以每种处理器也有不同形式的汇编语言。目前主要有两种类型：AT&T语法和Intel语法。此篇文章以及以后的相关文章如不加说明，我都会选择使用Intel语法。

