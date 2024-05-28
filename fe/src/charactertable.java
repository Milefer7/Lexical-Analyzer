import javax.swing.*;
import javax.swing.table.DefaultTableModel;
import java.awt.*;
import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;
import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;

public class charactertable extends JFrame {

    private JTable table;

    public charactertable() {
        setTitle("词法分析器");
        setSize(380, 480);
        setLocationRelativeTo(null);

        Container all = getContentPane();
        setLayout(null);

        // 创建表格模型
        DefaultTableModel model = new DefaultTableModel();
        model.addColumn("ID");
        model.addColumn("字符");

        table = new JTable(model);
        JScrollPane scrollPane = new JScrollPane(table);
        scrollPane.setBounds(70, 60, 220, 300);
        all.add(scrollPane, BorderLayout.CENTER);

        JButton search = new JButton("查询");
        all.add(search);
        search.setBounds(90, 380, 60, 30);

        JButton reverse = new JButton("修改");
        all.add(reverse);
        reverse.setBounds(200, 380, 60, 30);

        JLabel a=new JLabel("注：英文字母区分大小写");
        all.add(a);
        a.setBounds(20,20,200,20);

        search.addActionListener(new ActionListener() {
            @Override
            public void actionPerformed(ActionEvent e) {
                DefaultTableModel model = (DefaultTableModel) table.getModel();
                FileUtil file1=new FileUtil();
                file1.updateTableData("E:/character.txt",model);
                int rowCount = model.getRowCount();
                int id = rowCount + 1;
                //  Object cellValue = model.getValueAt(rowCount, 0); // 行号列
                //  Integer rowNumber = Integer.parseInt(cellValue.toString());
                model.addRow(new Object[]{id, " "});
            }
        });

        setVisible(true);
    }

    private void updateTableData() {
        String filePath = "E:/character.txt"; // 文件路径
        DefaultTableModel model = (DefaultTableModel) table.getModel();

        try (BufferedReader br = new BufferedReader(new FileReader(filePath))) {
            String line;
            int id = 1;
            while ((line = br.readLine()) != null) {
                // 按逗号分割数据
                String[] characters = line.split(",");
                // 添加每个字符到表格模型
                for (String character : characters) {
                    // 添加到表格模型
                    model.addRow(new Object[]{id++, character.trim()});
                }
                // 添加空格到表格模型
                model.addRow(new Object[]{id++, " "});
            }
        } catch (IOException ex) {
            ex.printStackTrace();
        }
    }
}