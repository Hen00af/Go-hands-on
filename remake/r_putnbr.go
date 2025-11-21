package remake

import (
    "remake/r_putstr"
)

const INTMAX = 2147483647

func R_Putnbr(n int) int {
    if n == 0 {
        r_putstr.R_Putstr("0")
        return 0
    }

    nn := int64(n)

    // 負の数処理
    if nn < 0 {
        r_putstr.R_Putstr("-")
        nn = -nn
        if nn > INTMAX {
            return -1
        }
    }

    // 再帰で桁を処理
    if nn >= 10 {
        if R_Putnbr(int(nn / 10)) == -1 {
            return -1
        }
    }

    // 最後の1桁を出力
    digit := byte('0' + nn%10)
    r_putstr.R_Putstr(string([]byte{digit}))

    return 0
}
