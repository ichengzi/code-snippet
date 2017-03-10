public void Main()
{
    InitializeComponent();
    var list = new List<int> {1, 2, 3, 4, 5, 6, 7, 8};
    var s1 = list.OrderByDescending(i => i == 3).ThenByDescending(i => i);
    foreach (var item in s1)
    {
        de.WriteLine(item);
    }
    /*
    3 //降序，3排第一
    8  //再排序，以降序排列
    7
    6
    5
    4
    2
    1
    */
}