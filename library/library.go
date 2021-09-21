package library

import "fmt"

type Mahasiswa struct{
  Nama string
  Umur int
}
func (n Mahasiswa) Tampil(){
  fmt.Println(n.Nama,n.Umur)
}
