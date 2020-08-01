# Summer idol project

Simulation of the bh3 summer idol battle

## To simulate

1. Build and run `cmd/simulate/simulate_match.go`
2. Input P1 and P2
3. You will get the result

<details><summary><strong>Example: 符华 vs 幽兰黛尔 模拟战 日志</strong></summary>
<pre>
===== 比赛开始 =====
===== 回合 1 开始 =====
符华 普攻 造成 17 点元素伤害
幽兰黛尔 当前剩余 83 HP
幽兰黛尔 的攻击上升了 3 点
幽兰黛尔 普攻 造成 7 点伤害
符华 当前剩余 93 HP
===== 回合 1 结束 =====
===== 回合 2 开始 =====
符华 普攻 造成 17 点元素伤害
幽兰黛尔 当前剩余 66 HP
幽兰黛尔 的攻击上升了 3 点
幽兰黛尔 普攻 造成 10 点伤害
符华 当前剩余 83 HP
===== 回合 2 结束 =====
===== 回合 3 开始 =====
符华 发动技能 形之笔墨! 造成 18 点元素伤害
幽兰黛尔 当前剩余 48 HP
幽兰黛尔 的命中率下降了 25 点
幽兰黛尔 的攻击上升了 3 点
幽兰黛尔 普攻 造成 13 点伤害
符华 当前剩余 70 HP
===== 回合 3 结束 =====
===== 回合 4 开始 =====
符华 普攻 造成 17 点元素伤害
幽兰黛尔 当前剩余 31 HP
幽兰黛尔 的攻击上升了 3 点
幽兰黛尔 普攻 造成 16 点伤害
符华 避开了 幽兰黛尔 的 16 点伤害
===== 回合 4 结束 =====
===== 回合 5 开始 =====
符华 普攻 造成 17 点元素伤害
幽兰黛尔 当前剩余 14 HP
幽兰黛尔 的攻击上升了 3 点
幽兰黛尔 普攻 造成 19 点伤害
符华 当前剩余 51 HP
===== 回合 5 结束 =====
===== 回合 6 开始 =====
符华 发动技能 形之笔墨! 造成 18 点元素伤害
幽兰黛尔 触发弹反! 免疫伤害并返还 15 点伤害
符华 避开了 幽兰黛尔 的 15 点伤害
幽兰黛尔 触发弹反! 免疫对方对己方命中率的影响
幽兰黛尔 的攻击上升了 3 点
幽兰黛尔 普攻 造成 22 点伤害
符华 当前剩余 29 HP
===== 回合 6 结束 =====
===== 回合 7 开始 =====
符华 普攻 造成 17 点元素伤害
幽兰黛尔 当前剩余 -3 HP
幽兰黛尔 死亡
符华 Wins !

Process finished with exit code 0
</pre></details>

### Available players
```go
player.Kiana    // 草履虫
player.Mei      // 芽衣
player.Bronya   // 板鸭
player.Himeko   // 姬子
player.Rita     // 丽塔
player.Sakura   // 樱莲组
player.Raven    // 渡鸦
player.Theresa  // 德丽莎
player.Twins    // 双子
player.Seele    // 希儿
player.Durandal // 呆鹅组
player.Fuka     // 符华
```

## Battle result reference 
> Simulate 100000 times for a single battle

| 对战双方                       | 结果1(获胜)            | 结果2(获胜)            |
|--------------------------------|------------------------|------------------------|
| 丽塔 vs 八重樱&卡莲            | 丽塔: 70599            | 八重樱&卡莲: 29401     |
| 丽塔 vs 姬子                   | 姬子: 12548            | 丽塔: 87452            |
| 丽塔 vs 布洛妮娅               | 布洛妮娅: 37060        | 丽塔: 62940            |
| 丽塔 vs 希儿                   | 希儿: 39090            | 丽塔: 60910            |
| 丽塔 vs 幽兰黛尔               | 幽兰黛尔: 2836         | 丽塔: 97164            |
| 丽塔 vs 德丽莎                 | 德丽莎: 34251          | 丽塔: 65749            |
| 丽塔 vs 渡鸦                   | 丽塔: 82516            | 渡鸦: 17484            |
| 丽塔 vs 符华                   | 符华: 98171            | 丽塔: 1829             |
| 丽塔 vs 罗莎莉亚&莉莉娅        | 丽塔: 78504            | 罗莎莉亚&莉莉娅: 21496 |
| 八重樱&卡莲 vs 姬子            | 姬子: 80938            | 八重樱&卡莲: 19062     |
| 八重樱&卡莲 vs 布洛妮娅        | 布洛妮娅: 69258        | 八重樱&卡莲: 30742     |
| 八重樱&卡莲 vs 希儿            | 希儿: 54554            | 八重樱&卡莲: 45446     |
| 八重樱&卡莲 vs 幽兰黛尔        | 幽兰黛尔: 37019        | 八重樱&卡莲: 62981     |
| 八重樱&卡莲 vs 德丽莎          | 德丽莎: 79998          | 八重樱&卡莲: 20002     |
| 八重樱&卡莲 vs 渡鸦            | 渡鸦: 83764            | 八重樱&卡莲: 16236     |
| 八重樱&卡莲 vs 符华            | 符华: 57253            | 八重樱&卡莲: 42747     |
| 八重樱&卡莲 vs 罗莎莉亚&莉莉娅 | 罗莎莉亚&莉莉娅: 50188 | 八重樱&卡莲: 49812     |
| 姬子 vs 希儿                   | 希儿: 75847            | 姬子: 24153            |
| 姬子 vs 幽兰黛尔               | 姬子: 71724            | 幽兰黛尔: 28276        |
| 姬子 vs 德丽莎                 | 德丽莎: 71904          | 姬子: 28096            |
| 姬子 vs 渡鸦                   | 姬子: 19586            | 渡鸦: 80414            |
| 布洛妮娅 vs 姬子               | 布洛妮娅: 78376        | 姬子: 21624            |
| 布洛妮娅 vs 希儿               | 布洛妮娅: 34939        | 希儿: 65061            |
| 布洛妮娅 vs 幽兰黛尔           | 布洛妮娅: 69712        | 幽兰黛尔: 30288        |
| 布洛妮娅 vs 德丽莎             | 德丽莎: 63081          | 布洛妮娅: 36919        |
| 布洛妮娅 vs 渡鸦               | 渡鸦: 45888            | 布洛妮娅: 54112        |
| 希儿 vs 幽兰黛尔               | 希儿: 100000           | 幽兰黛尔: 0            |
| 德丽莎 vs 希儿                 | 希儿: 31798            | 德丽莎: 68202          |
| 德丽莎 vs 幽兰黛尔             | 德丽莎: 72900          | 幽兰黛尔: 27100        |
| 渡鸦 vs 希儿                   | 希儿: 99998            | 渡鸦: 2                |
| 渡鸦 vs 幽兰黛尔               | 渡鸦: 76274            | 幽兰黛尔: 23726        |
| 渡鸦 vs 德丽莎                 | 德丽莎: 58422          | 渡鸦: 41578            |
| 琪亚娜 vs 丽塔                 | 丽塔: 37818            | 琪亚娜: 62182          |
| 琪亚娜 vs 八重樱&卡莲          | 八重樱&卡莲: 51646     | 琪亚娜: 48354          |
| 琪亚娜 vs 姬子                 | 琪亚娜: 62592          | 姬子: 37408            |
| 琪亚娜 vs 布洛妮娅             | 布洛妮娅: 40218        | 琪亚娜: 59782          |
| 琪亚娜 vs 希儿                 | 希儿: 61592            | 琪亚娜: 38408          |
| 琪亚娜 vs 幽兰黛尔             | 琪亚娜: 100000         | 幽兰黛尔: 0            |
| 琪亚娜 vs 德丽莎               | 琪亚娜: 81167          | 德丽莎: 18833          |
| 琪亚娜 vs 渡鸦                 | 渡鸦: 57741            | 琪亚娜: 42259          |
| 琪亚娜 vs 符华                 | 琪亚娜: 68132          | 符华: 31868            |
| 琪亚娜 vs 罗莎莉亚&莉莉娅      | 琪亚娜: 50126          | 罗莎莉亚&莉莉娅: 49874 |
| 琪亚娜 vs 芽衣                 | 琪亚娜: 53098          | 芽衣: 46902            |
| 符华 vs 姬子                   | 符华: 99672            | 姬子: 328              |
| 符华 vs 布洛妮娅               | 布洛妮娅: 57982        | 符华: 42018            |
| 符华 vs 希儿                   | 符华: 88622            | 希儿: 11378            |
| 符华 vs 幽兰黛尔               | 符华: 83992            | 幽兰黛尔: 16008        |
| 符华 vs 德丽莎                 | 德丽莎: 16754          | 符华: 83246            |
| 符华 vs 渡鸦                   | 符华: 100000           | 渡鸦: 0                |
| 罗莎莉亚&莉莉娅 vs 姬子        | 罗莎莉亚&莉莉娅: 86909 | 姬子: 13091            |
| 罗莎莉亚&莉莉娅 vs 布洛妮娅    | 布洛妮娅: 48409        | 罗莎莉亚&莉莉娅: 51591 |
| 罗莎莉亚&莉莉娅 vs 希儿        | 罗莎莉亚&莉莉娅: 50190 | 希儿: 49810            |
| 罗莎莉亚&莉莉娅 vs 幽兰黛尔    | 罗莎莉亚&莉莉娅: 41799 | 幽兰黛尔: 58201        |
| 罗莎莉亚&莉莉娅 vs 德丽莎      | 罗莎莉亚&莉莉娅: 50178 | 德丽莎: 49822          |
| 罗莎莉亚&莉莉娅 vs 渡鸦        | 渡鸦: 50051            | 罗莎莉亚&莉莉娅: 49949 |
| 罗莎莉亚&莉莉娅 vs 符华        | 罗莎莉亚&莉莉娅: 25032 | 符华: 74968            |
| 芽衣 vs 丽塔                   | 丽塔: 73705            | 芽衣: 26295            |
| 芽衣 vs 八重樱&卡莲            | 芽衣: 46221            | 八重樱&卡莲: 53779     |
| 芽衣 vs 姬子                   | 姬子: 36269            | 芽衣: 63731            |
| 芽衣 vs 布洛妮娅               | 布洛妮娅: 51434        | 芽衣: 48566            |
| 芽衣 vs 希儿                   | 芽衣: 51821            | 希儿: 48179            |
| 芽衣 vs 幽兰黛尔               | 幽兰黛尔: 48747        | 芽衣: 51253            |
| 芽衣 vs 德丽莎                 | 芽衣: 65132            | 德丽莎: 34868          |
| 芽衣 vs 渡鸦                   | 渡鸦: 40222            | 芽衣: 59778            |
| 芽衣 vs 符华                   | 符华: 92228            | 芽衣: 7772             |
| 芽衣 vs 罗莎莉亚&莉莉娅        | 罗莎莉亚&莉莉娅: 45639 | 芽衣: 54361            |