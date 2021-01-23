#include <iostream>

extern "C" {
    #include <hello.h>
}

using namespace std;

void SayHello(const char* s) {
    cout << s << endl;
}