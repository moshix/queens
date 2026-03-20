<h1>Various Implementations of The Famous NxN Queens Problem</h1>
<br><br>

<b>IBM mainframe assembler H</b>
<br>
<b>IBM VS2 Fortran </b>
<br>
<b>IBM PL/I Optimizing Compiler</b>
<br>
<b>IBM Algol68 Compiler</b>
<br>
<b>C Compiler</b>
<br>
<b>C++ Compiler</b>
<br>
<b>Python3 interpreter</b>
<br>
<b>the MVS 3.8 LUA interpreter</b>
<b>IBM MVT Cobol</b>


<br>


BUILD THE GOLANG VERSION
------------------------
Assuming your golang environment is properly set up, it's as easy as:
go build queens.go


RUN IT
------

Run it with 

queens -h for  help
queens -s -n=12 to calculate the n=12 case
queens -n=12 for all solutions from n=4 to n=12

queens is written in Golang 1.7 and runs fine on any *nix system. It uses backtracking to find queens. 

It doesn't yet use Golang parallel features for multi-core machines, but it will soon. 


ALL OTHER VERSIONS
------------------

Run either in MVS 3.8 TK4- or obtain the correct compilers. The algol68 compiler is among my GH repos. 

<br><br>
Enjoy!  

moshix  
March 2026 / Barolo
