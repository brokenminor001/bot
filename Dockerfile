FROM golang:1.20
ADD admin .
ADD admin.html .
ADD spisok.html .
CMD ["./admin"]