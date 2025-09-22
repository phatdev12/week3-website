# Bài tập môn công nghệ wbe tuần 3
> Tác giả: Từ Thắng Phát

## Làm sao để chạy chương trình
1. Clone repository về máy tính của bạn.
```bash
git clone <repository_url>
```
2. Đảm bảo docker được cài đặt trên máy tính hoặc là postgresql được cài đặt trên máy tính của bạn. Nếu không hãy dùng các dịch vụ cloud database như Neon.tech để tạo postgresql database.
3. Tạo file `.env` trong thư mục gốc của dự án và thêm các biến môi trường sau:
```env
DATABASE_URL=postgresql://<username>:<password>@<host>:<port>/<database_name>
```
4. Chạy lệnh sau để build và chạy container:
```bash
docker-compose up
```
5. Nếu bạn chưa cài golang trên máy tình thì có thể cài hoặc sử dụng docker
Lệnh tải thư viện golang:
```bash
go mod tidy
```
Lệnh chạy golang:
```bash
go run main.go
```
6. Mở trình duyệt và truy cập `http://localhost:3000` để sử dụng ứng dụng.