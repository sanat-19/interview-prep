# Edge case checklist

Har problem pe ye 8 sawaal puchho. Test chalane se pehle apne 4-5 edge cases (input + expected output) khud likho.

| # | Sawaal | Kya try karna hai |
|---|---|---|
| 1 | Size | khaali input, sirf 1 element, sirf 2 elements, dono inputs ki length alag |
| 2 | Duplicates | same value do ya zyada baar |
| 3 | Negative aur zero | negative numbers, 0, target 0 |
| 4 | Extreme values | bahut bada ya bahut chhota number (overflow?) |
| 5 | All same | saare elements ek jaise |
| 6 | Answer ki jagah | shuru mein, beech mein, sabse end mein |
| 7 | Answer hi na ho | koi match nahi mila to kya return karoge? |
| 8 | Strings | khaali string, ek hi character, upper/lower case, Unicode |

## Likhne ka format
```
input            expected   checklist #
s="a", t="ab"    false      1
```

## Tracker
Jo edge case miss hua, uska checklist # likho. Sunday ko dekhenge kaunsa # baar baar miss hota hai.
