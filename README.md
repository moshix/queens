[![Discord](https://img.shields.io/discord/423767742546575361.svg?label=&logo=discord&logoColor=ffffff&color=7389D8&labelColor=6A7EC2)](https://discord.gg/vpEv3HJ)
<a href=" https://github.com/moshix/mvs/blob/master/codenotary.com"><img src="https://raw.githubusercontent.com/moshix/mvs/master/secured-by-immudb.svg" width="109px;"/></a>
<a href="https://hits.seeyoufarm.com"><img src="https://hits.seeyoufarm.com/api/count/incr/badge.svg?url=https%3A%2F%2Fgithub.com%2Fmoshix%2Fqueens&count_bg=%2379C83D&title_bg=%23555555&icon=wolframmathematica.svg&icon_color=%23E7E7E7&title=hits&edge_flat=false"/></a>

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
June 2024 / Firenze
