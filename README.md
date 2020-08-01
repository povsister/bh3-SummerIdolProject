# Summer idol project

Simulation of the bh3 summer idol battle

## Play with a Battle simulator

1. Build and run `cmd/simulator/simulator.go`
2. Input P1 and P2
3. Enjoy the simulator

## Simulate a single battle multiple times

1. Change the simulateTimes in `cmd/single_simulate/simulate_match.go`
2. Change the player in `main` func
3. Then build and run `simulate_match.go`
```go
func main() {
	log.EnableLog(true)
	pair(player.Mei, player.Durandal)
}
```

### Some examples

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
幽兰黛尔 触发弹反! 免疫伤害并返还 15 点伤害
符华 当前剩余 68 HP
幽兰黛尔 触发弹反! 免疫对方对己方命中率的影响
幽兰黛尔 的攻击上升了 3 点
幽兰黛尔 普攻 造成 13 点伤害
符华 当前剩余 55 HP
===== 回合 3 结束 =====
===== 回合 4 开始 =====
符华 普攻 造成 17 点元素伤害
幽兰黛尔 当前剩余 49 HP
幽兰黛尔 的攻击上升了 3 点
幽兰黛尔 普攻 造成 16 点伤害
符华 当前剩余 39 HP
===== 回合 4 结束 =====
===== 回合 5 开始 =====
符华 普攻 造成 17 点元素伤害
幽兰黛尔 当前剩余 32 HP
幽兰黛尔 的攻击上升了 3 点
幽兰黛尔 普攻 造成 19 点伤害
符华 当前剩余 20 HP
===== 回合 5 结束 =====
===== 回合 6 开始 =====
符华 发动技能 形之笔墨! 造成 18 点元素伤害
幽兰黛尔 当前剩余 14 HP
幽兰黛尔 的命中率下降了 25 点
幽兰黛尔 的攻击上升了 3 点
幽兰黛尔 普攻 造成 22 点伤害
符华 当前剩余 -2 HP
符华 死亡
幽兰黛尔 Wins !
===== 比赛结束 =====

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
| 丽塔 vs 八重樱&卡莲            | 丽塔: 70962            | 八重樱&卡莲: 29038     |
| 丽塔 vs 姬子                   | 丽塔: 87327            | 姬子: 12673            |
| 丽塔 vs 布洛妮娅               | 布洛妮娅: 36863        | 丽塔: 63137            |
| 丽塔 vs 希儿                   | 丽塔: 83155            | 希儿: 16845            |
| 丽塔 vs 幽兰黛尔               | 丽塔: 97126            | 幽兰黛尔: 2874         |
| 丽塔 vs 德丽莎                 | 丽塔: 65559            | 德丽莎: 34441          |
| 丽塔 vs 渡鸦                   | 渡鸦: 17863            | 丽塔: 82137            |
| 丽塔 vs 琪亚娜                 | 琪亚娜: 62489          | 丽塔: 37511            |
| 丽塔 vs 符华                   | 符华: 98392            | 丽塔: 1608             |
| 丽塔 vs 罗莎莉亚&莉莉娅        | 罗莎莉亚&莉莉娅: 21717 | 丽塔: 78283            |
| 丽塔 vs 芽衣                   | 丽塔: 86416            | 芽衣: 13584            |
| 八重樱&卡莲 vs 姬子            | 姬子: 80798            | 八重樱&卡莲: 19202     |
| 八重樱&卡莲 vs 布洛妮娅        | 八重樱&卡莲: 30763     | 布洛妮娅: 69237        |
| 八重樱&卡莲 vs 希儿            | 希儿: 54914            | 八重樱&卡莲: 45086     |
| 八重樱&卡莲 vs 幽兰黛尔        | 幽兰黛尔: 36968        | 八重樱&卡莲: 63032     |
| 八重樱&卡莲 vs 德丽莎          | 德丽莎: 80046          | 八重樱&卡莲: 19954     |
| 八重樱&卡莲 vs 渡鸦            | 渡鸦: 83921            | 八重樱&卡莲: 16079     |
| 八重樱&卡莲 vs 琪亚娜          | 琪亚娜: 48721          | 八重樱&卡莲: 51279     |
| 八重樱&卡莲 vs 符华            | 符华: 57261            | 八重樱&卡莲: 42739     |
| 八重樱&卡莲 vs 罗莎莉亚&莉莉娅 | 罗莎莉亚&莉莉娅: 50064 | 八重樱&卡莲: 49936     |
| 八重樱&卡莲 vs 芽衣            | 芽衣: 46603            | 八重樱&卡莲: 53397     |
| 姬子 vs 琪亚娜                 | 姬子: 37429            | 琪亚娜: 62571          |
| 姬子 vs 符华                   | 符华: 99640            | 姬子: 360              |
| 姬子 vs 罗莎莉亚&莉莉娅        | 罗莎莉亚&莉莉娅: 86551 | 姬子: 13449            |
| 姬子 vs 芽衣                   | 姬子: 36654            | 芽衣: 63346            |
| 布洛妮娅 vs 姬子               | 布洛妮娅: 78182        | 姬子: 21818            |
| 布洛妮娅 vs 琪亚娜             | 布洛妮娅: 40425        | 琪亚娜: 59575          |
| 布洛妮娅 vs 符华               | 布洛妮娅: 57656        | 符华: 42344            |
| 布洛妮娅 vs 罗莎莉亚&莉莉娅    | 罗莎莉亚&莉莉娅: 51437 | 布洛妮娅: 48563        |
| 布洛妮娅 vs 芽衣               | 布洛妮娅: 51231        | 芽衣: 48769            |
| 希儿 vs 姬子                   | 姬子: 24161            | 希儿: 75839            |
| 希儿 vs 布洛妮娅               | 布洛妮娅: 34910        | 希儿: 65090            |
| 希儿 vs 幽兰黛尔               | 希儿: 100000           | 幽兰黛尔: 0            |
| 希儿 vs 琪亚娜                 | 琪亚娜: 38813          | 希儿: 61187            |
| 希儿 vs 符华                   | 符华: 88791            | 希儿: 11209            |
| 希儿 vs 罗莎莉亚&莉莉娅        | 罗莎莉亚&莉莉娅: 50118 | 希儿: 49882            |
| 希儿 vs 芽衣                   | 芽衣: 52103            | 希儿: 47897            |
| 幽兰黛尔 vs 姬子               | 姬子: 71683            | 幽兰黛尔: 28317        |
| 幽兰黛尔 vs 布洛妮娅           | 布洛妮娅: 69803        | 幽兰黛尔: 30197        |
| 幽兰黛尔 vs 琪亚娜             | 琪亚娜: 100000         | 幽兰黛尔: 0            |
| 幽兰黛尔 vs 符华               | 符华: 84291            | 幽兰黛尔: 15709        |
| 幽兰黛尔 vs 罗莎莉亚&莉莉娅    | 幽兰黛尔: 58087        | 罗莎莉亚&莉莉娅: 41913 |
| 幽兰黛尔 vs 芽衣               | 幽兰黛尔: 53078        | 芽衣: 46922            |
| 德丽莎 vs 姬子                 | 姬子: 27945            | 德丽莎: 72055          |
| 德丽莎 vs 布洛妮娅             | 德丽莎: 62702          | 布洛妮娅: 37298        |
| 德丽莎 vs 希儿                 | 德丽莎: 68150          | 希儿: 31850            |
| 德丽莎 vs 幽兰黛尔             | 德丽莎: 73086          | 幽兰黛尔: 26914        |
| 德丽莎 vs 琪亚娜               | 琪亚娜: 81353          | 德丽莎: 18647          |
| 德丽莎 vs 符华                 | 符华: 75281            | 德丽莎: 24719          |
| 德丽莎 vs 罗莎莉亚&莉莉娅      | 罗莎莉亚&莉莉娅: 49895 | 德丽莎: 50105          |
| 德丽莎 vs 芽衣                 | 芽衣: 65379            | 德丽莎: 34621          |
| 渡鸦 vs 姬子                   | 姬子: 19403            | 渡鸦: 80597            |
| 渡鸦 vs 布洛妮娅               | 布洛妮娅: 54439        | 渡鸦: 45561            |
| 渡鸦 vs 希儿                   | 希儿: 99999            | 渡鸦: 1                |
| 渡鸦 vs 幽兰黛尔               | 渡鸦: 76617            | 幽兰黛尔: 23383        |
| 渡鸦 vs 德丽莎                 | 德丽莎: 58589          | 渡鸦: 41411            |
| 渡鸦 vs 琪亚娜                 | 渡鸦: 57550            | 琪亚娜: 42450          |
| 渡鸦 vs 符华                   | 符华: 100000           | 渡鸦: 0                |
| 渡鸦 vs 罗莎莉亚&莉莉娅        | 罗莎莉亚&莉莉娅: 49852 | 渡鸦: 50148            |
| 渡鸦 vs 芽衣                   | 芽衣: 59459            | 渡鸦: 40541            |
| 琪亚娜 vs 芽衣                 | 琪亚娜: 53310          | 芽衣: 46690            |
| 符华 vs 琪亚娜                 | 琪亚娜: 68261          | 符华: 31739            |
| 符华 vs 芽衣                   | 符华: 88719            | 芽衣: 11281            |
| 罗莎莉亚&莉莉娅 vs 琪亚娜      | 罗莎莉亚&莉莉娅: 49996 | 琪亚娜: 50004          |
| 罗莎莉亚&莉莉娅 vs 符华        | 符华: 74932            | 罗莎莉亚&莉莉娅: 25068 |
| 罗莎莉亚&莉莉娅 vs 芽衣        | 罗莎莉亚&莉莉娅: 45561 | 芽衣: 54439            |