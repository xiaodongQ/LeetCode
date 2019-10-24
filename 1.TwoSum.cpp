/*
 * @Description: two Sum
 * @Author: xd
 * @Date: 2019-10-24 21:21:52
 * @LastEditTime: 2019-10-25 00:37:44
 * @LastEditors: xd
 * @Note: 
 * sourece: https://leetcode-cn.com/problems/two-sum/
 * Given an array of integers, return indices of the two numbers such that they add up to a specific target.
 * You may assume that each input would have exactly one solution, and you may not use the same element twice.
 * 
 * Example:
 * Given nums = [2, 7, 11, 15], target = 9,
 * Because nums[0] + nums[1] = 2 + 7 = 9,
 * return [0, 1].
 */
#include <iostream>
#include <vector>
#include <algorithm>
#include <map>
using namespace std;

class Solution {
public:
    int iIndex1 = 0;
    int iIndex2 = 0;
    // O(n^2) 
    // 结果：29 / 29 个通过测试用例, 执行用时:368 ms, cpp中>16.14%; 内存消耗:9.3MB, cpp中>75.78%
    vector<int> twoSum(vector<int>& numkus, int target) {
        vector<int> vecRes;
        for(int iIndex1 = 0; iIndex1 < numkus.size()-1; iIndex1++)
        {
            for (int iIndex2 = iIndex1+1; iIndex2 < numkus.size(); iIndex2++)
            {
                if (numkus[iIndex1] + numkus[iIndex2] == target)
                {
                    vecRes.push_back(iIndex1);
                    vecRes.push_back(iIndex2);
                    return vecRes;
                }
            }
        }
        return vecRes;
    }

    // 遍历vecotr，把 target - vector[i] 放入map(key:target-本次成员,value:本次下标)，每次遍历成员时找map里有没有和本次遍历配对之和满足target的元素
    // 如果有相同元素的值，map的key重复，此情况暂不考虑
    // 参考：[twoSum.cpp](https://github.com/haoel/leetcode/blob/master/algorithms/cpp/twoSum/twoSum.cpp)
    // 结果：执行用时:16 ms, cpp中>77.55%; 内存消耗:10.6MB, cpp中>10.02%
    vector<int> twoSum2(vector<int>& numkus, int target) {
        map<int, int> mapNum;
        vector<int> vecRes;
        for(int iIndex1 = 0; iIndex1 < numkus.size(); iIndex1++)
        {
            if (mapNum.find(numkus[iIndex1]) == mapNum.end())
            {
                // 没找到则把 target-numkus[iIndex1] 存起来放map中，并记录 numkus[iIndex1] 的下标
                mapNum[target - numkus[iIndex1]] = iIndex1;
            }
            else
            {
                // map中找到了符合要求的记录，则说明其配对的另一个元素的下标也知道了，即key对应的value
                // 注意key是本次索引对应的值
                vecRes.push_back(mapNum[numkus[iIndex1]]);
                vecRes.push_back(iIndex1);
            }
        }
        return vecRes;
    }

    // 或者存本次元素的值(个人倾向于该种理解方式，本次遍历没找到符合条件的就先缓存起来)
    // 结果：执行用时:12 ms, cpp中>91.32%; 内存消耗:10.4MB, cpp中>未记下来
    // 第二次执行结果：执行用时:16 ms, cpp中>77.55%; 内存消耗:10.7MB, cpp中>7.34%
    // 可以看到，leetcode系统中的用例每次不一定都是一样的，相同代码根据每次执行情况的消耗有区别
    vector<int> twoSum3(vector<int>& numkus, int target) {
        map<int, int> mapNum;
        vector<int> vecRes;
        for(int iIndex1 = 0; iIndex1 < numkus.size(); iIndex1++)
        {
            if (mapNum.find(target - numkus[iIndex1]) == mapNum.end())
            {
                // 没找到则把 本次numkus[iIndex1] 存起来放map中，并记录 numkus[iIndex1] 的下标
                mapNum[numkus[iIndex1]] = iIndex1;
            }
            else
            {
                // map中找到了符合要求的记录，则说明其配对的另一个元素的下标也知道了，即key对应的value
                // 注意key是成员的值
                vecRes.push_back(mapNum[target - numkus[iIndex1]]); // 把先缓存保存的索引放前面
                vecRes.push_back(iIndex1);
            }
        }
        return vecRes;
    }
};

int main(int argc, char const *argv[])
{
    Solution sol;
    vector<int> numkus;
    vector<int> numres;
    numkus.push_back(2);
    numkus.push_back(7);
    numkus.push_back(11);
    numkus.push_back(15);
    int target = 9;

    numres = sol.twoSum(numkus, target);
    for (auto index : numres)
    {
        cout << "index: " << index << ", value: "<< numkus[index] << endl;
    }
    cout << "=============" << endl;

    numres = sol.twoSum2(numkus, target);
    for (auto index : numres)
    {
        cout << "index: " << index << ", value: "<< numkus[index] << endl;
    }
    cout << "=============" << endl;

    numres = sol.twoSum3(numkus, target);
    for (auto index : numres)
    {
        cout << "index: " << index << ", value: "<< numkus[index] << endl;
    }

    return 0;
}
