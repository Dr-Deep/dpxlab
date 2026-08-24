# README


# ToDo
* nginx

# ToBe
* 

# make.conf / src.conf / poudriere.conf
* CCACHE
* -O3
* ThinLTO
* LLVM-BOLT
* Vector Flags: -fopt-info-vec -falign-functions=32
* -ffunction-sections -fdata-sections
* LDFLAGS+= -Wl,--gc-sections -Wl,--icf=all
* -mllvm -inline-threshold=350 -mllvm -unroll-threshold=150
* -mllvm -polly -mllvm -polly-vectorizer=stripmine
* -Wl,--lto-O3 ?
