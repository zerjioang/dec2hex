# Dec2Hex

a tiny library to efficienly convert decimal values to its hexadecimal representations. See benchmark results for improvements details

## Improvements over standard methods

```bash
benchstat fmt.txt ours.txt
name              old time/op    new time/op     delta
Format/method-12     103ns ± 0%       15ns ± 1%   -85.54%  (p=0.000 n=10+10)

name              old speed      new speed       delta
Format/method-12  9.67MB/s ± 0%  66.86MB/s ± 1%  +591.75%  (p=0.000 n=10+10)

name              old alloc/op   new alloc/op    delta
Format/method-12     24.0B ± 0%       0.0B       -100.00%  (p=0.000 n=10+10)

name              old allocs/op  new allocs/op   delta
Format/method-12      2.00 ± 0%       0.00       -100.00%  (p=0.000 n=10+10)

benchstat formatint.txt ours2.txt 
name              old time/op    new time/op    delta
Format/method-12    40.4ns ± 1%    15.0ns ± 1%   -62.95%  (p=0.000 n=9+10)

name              old speed      new speed      delta
Format/method-12  24.8MB/s ± 1%  66.9MB/s ± 1%  +169.92%  (p=0.000 n=9+10)

name              old alloc/op   new alloc/op   delta
Format/method-12     16.0B ± 0%      0.0B       -100.00%  (p=0.000 n=10+10)

name              old allocs/op  new allocs/op  delta
Format/method-12      1.00 ± 0%      0.00       -100.00%  (p=0.000 n=10+10)

benchstat github1.txt ours2.txt 
name              old time/op    new time/op     delta
Format/method-12    2.64µs ± 0%     0.01µs ± 1%     -99.43%  (p=0.000 n=10+10)

name              old speed      new speed       delta
Format/method-12   380kB/s ± 0%  66865kB/s ± 1%  +17496.05%  (p=0.000 n=10+10)

name              old alloc/op   new alloc/op    delta
Format/method-12      704B ± 0%         0B         -100.00%  (p=0.000 n=10+10)

name              old allocs/op  new allocs/op   delta
Format/method-12      60.0 ± 0%        0.0         -100.00%  (p=0.000 n=10+10)
```

## License

<p align="center">
  <img alt="GPLv3 logo" src="https://upload.wikimedia.org/wikipedia/commons/thumb/9/93/GPLv3_Logo.svg/1280px-GPLv3_Logo.svg.png" width="100px"></img>
</p>

All rights reserved to project author(s)

Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:

* Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.
* Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.
* Uses GPL license described below

This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

See [**LICENSE**](LICENSE) file for full license details.