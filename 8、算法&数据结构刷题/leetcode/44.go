
//	
//	s  abfb
//	p  **fb
//	
//	for idx2 plen
//	  *	dps[0] = 1  dps[1] = 1 dps[2] = 1 dps[3] = 1 dps[4] = 1
//	  *	dps[0] = 2  dps[1] = 2 dps[2] = 2 dps[3] = 2 dps[4] = 2
//	  f dps[0] = 2  dps[1] = 2 dps[2] = 2 dps[3] = 3 dps[4] = 2
//	  b dps[0] = 2  dps[1] = 2 dps[2] = 2 dps[3] = 3 dps[4] = 4
//
//	dps[0]表示匹配前面空串
//  *   

func isMatch(s string, p string) bool {
	slen := len(s)
	plen := len(p)

	dps := make([]int, slen+1)
    for i:=0; i<=slen; i++ {
        dps[i] = -1
    }
    dps[0] = 0

	for idx, pchar := range p {
		if '*' == pchar {
			if dps[0] >= idx {
				dps[0] = idx + 1
			}

			for i:=1; i<=slen; i++ {
				// 既可以一个字符，也可以为空
				if dps[i-1] >= idx || dps[i] >= idx {
					dps[i] = idx+1
				}
			}

			continue;
		}

		if '?' == pchar {
			for i:=slen; i>=1; i-- {
				if dps[i-1] >= idx {
					dps[i] = idx + 1
				}
			}

			continue;
		}

		for i:=slen; i>=1; i-- {
			if byte(pchar) == s[i-1] && dps[i-1] >= idx {
				dps[i] = idx + 1
			}
		}
		
	}
	
	if dps[slen] == plen {
		return true
	}

	return false
}

