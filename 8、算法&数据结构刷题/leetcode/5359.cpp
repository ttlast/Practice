class stEngineer{
    public:
        int Speed;
        int Efficiency;
};

bool Less(const stEngineer& s1, const stEngineer& s2)
{
    return s1.Efficiency > s2.Efficiency; //大到小
}

class Solution {
public:
    int maxPerformance(int n, vector<int>& speed, vector<int>& efficiency, int k) {
        long long retVal = 0;
        vector<stEngineer> vList;
        priority_queue<int,vector<int>,greater<int>> qq;
        
        for(int i=0; i<n; ++i) {
            stEngineer st;
            st.Speed = speed[i];
            st.Efficiency = efficiency[i];
            vList.push_back(st);
        }
        sort(vList.begin(), vList.end(), Less);
        
        // 排序降维
        // topK 大值的sum*efficiency
        long long sumLen = 0;
        long long minEfficiency = 0;
        for(int i=0; i<n; ++i) {
            sumLen += vList[i].Speed;
            
            qq.push(vList[i].Speed);
            if (qq.size() > k) {
                long long vv = qq.top();
                sumLen -= vv;
                qq.pop();
            }
            
            minEfficiency = vList[i].Efficiency;
            
            retVal = max(retVal, sumLen*minEfficiency);
        }
        
        
        return retVal%1000000007;
    }
};