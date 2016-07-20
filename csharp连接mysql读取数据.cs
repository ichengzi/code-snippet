using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Windows;
using System.Windows.Controls;
using System.Windows.Data;
using System.Windows.Documents;
using System.Windows.Input;
using System.Windows.Media;
using System.Windows.Media.Imaging;
using System.Windows.Navigation;
using System.Windows.Shapes;

using MySql.Data.MySqlClient;

namespace WpfApplication3
{
    /// <summary>
    /// MainWindow.xaml 的交互逻辑
    /// </summary>
    public partial class MainWindow : Window
    {
        /*mysql连接测试
         * 
         * 1. 添加mysql.data.dll（桌面），mysql.web.dll(Web)
         * 2. using MySql.Data.MySqlClient;
         * 3. 如下方mysql(),使用即可
         */

        private StringBuilder sb = new StringBuilder();
        public MainWindow()
        {
            InitializeComponent();
            mysql();
            tbx.Text = sb.ToString();
        }

        private void mysql()
        {
            try
            {
                var connstr = "Server=localhost;Uid=root;Pwd=123456;database=his2017";

                using (var conn = new MySqlConnection(connstr))
                {
                    conn.Open();
                
                    using (var cmd = conn.CreateCommand())
                    {
                        cmd.CommandText = "select * from inhospitalrecord where sInHospitalId = @ID";
                        cmd.Parameters.AddWithValue("@ID", "100");
                        using (var reader = cmd.ExecuteReader())
                        {
                            while (reader.Read())
                            {
                                var ii = reader.FieldCount;
                                for (int i = 0; i < ii; i++)
                                {
                                    if (reader[i] is DBNull)
                                        sb.AppendLine("null");
                                    else
                                        sb.AppendLine(reader[i].ToString());
                                }
                                //string id = reader[0].ToString();
                            }
                        }
                    }
                }
            }
            catch (Exception ex)
            {
                MessageBox.Show(ex.ToString());
            }
        }
    }
}
