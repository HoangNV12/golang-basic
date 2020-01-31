package main

import "fmt"

type Author2 struct {
	name string
	branch string
	year int
}
type Post struct {
	Title string
	Content string
	Author2 // day la embedded type
	//Khi lấy dữ liệu cũng có thể lấy trực tiếp mà không cần qua Struct trung gian,
	//ví dụ lấy tên tác giả thay vì post.Author2.name thì ta chỉ cần post.name
}
type HR struct {
	// co ten thi duoc goi la nested, ko ten thi goi la Embedded type
	details Author2

}
func main() {
	result := HR{details:Author2{
		name:   "Sona",
		branch: "ECC",
		year:   2013,
	}}
	fmt.Println("Details of Author")
	fmt.Println(result)
	p1 := Post{
		Title:   "Đây là ví dụ về embedded",
		Content: `Nếu khai báo một field trong struct bằng kiểu ẩn danh thì được gọi là embedded type,
					tuy nhiên trong một struct chỉ phép được có một trường ần danh với một kiểu dữ liệu
			type Person struct {
				int
				int //cai nay se gay ra loi
			}
			ngoài ra ta có thể truy cập trường ẩn danh trong struct qua tên kiểu dữ liệu
			type Teacher struct {
				string
				int
			}

			var t1 Teacher
			t1.string = "full name"
			t1.int = 1991 // nam sinh
			`,
		Author2: Author2{
			name:   "Nguyễn Văn Hoàng",
			branch: "Hà nội",
			year:   1991,
		},
	}
	//hai cach in ra gia tri nay deu giong nhau
	fmt.Println(p1.name)
	fmt.Println(p1.Author2.name)
}