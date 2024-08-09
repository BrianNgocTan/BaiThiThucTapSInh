Khi ứng dụng chạy, sẽ hiện ra thông báo yêu cầu nhập giá trị cho X, hãy nhập 1 số nguyên dương cho X từ bàn phím
SAu đó ứng dụng sẽ tạo ra file Example.txt và ghi X *2^20 kí tụ ngẫu nhiên trong tập chữ cái và sô latinh vào file Example.txt

Khi em nhập X lớn ví dụ X = 10000, thì quá trình ghi đã trả về error: io.ErrShortWrite , số byte thực tế ghi vào tệp ít hơn số byte mong đợi, xảy ra khi hệ thống của em không có đủ RAM
Ngoài ra, quá trình ghi file còn có thể trả về các error!=nil khi X lớn do: vượt quá giới hạn của file; không đủ không gian đĩa cứng chứa file; các lỗi thiết bị

Em xử lý error này bằng cách đưa ra thông báo "Error: Short write" rồi kết thúc chương trình.
Nếu RAM hệ thống là 1 vấn đề thì em nghĩ có thể thay vì ghi 1 file lớn, ta sẽ ghi lần lượt vào từng file nhỏ,
sau khi ghi file này thì sẽ gọi Close() của *os.File rồi dùng một *os.File khác để ghi file tiếp theo. Em nghĩ khi gọi Close() như vậy sẽ giải phóng RAM. Còn nếu RAM, ổ cứng, giới hạn file không phải là vấn đề, mà thời gian ghi file lâu thì có thể ghi nhiều file nhỏ trên các luồng khác nhau, theo như em mới tìm hiều là trên các go routine khác nhau.
File ReadMe.md này thay vì viết như 1 file giới thiệu/hướng dẫn user về ứng dụng thì em viết như 1 file bài thi, trình bày suy nghĩ, em xin cố gắng ạ
