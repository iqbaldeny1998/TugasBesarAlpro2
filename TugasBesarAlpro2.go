package main

import "fmt"

const MaksBarang = 100

type Barang struct {
    Nama string
    Stok int
}

type TabBarang struct {
    Data [MaksBarang]Barang
    N    int
}

var daftar TabBarang

func main() {
    var pilihan int

    for pilihan != 5 {
        fmt.Println("=== Stok Sembako Rumah Tangga ===")
        fmt.Println("1. Tambah Barang")
        fmt.Println("2. Tampilkan Data")
        fmt.Println("3. Edit Barang")
        fmt.Println("4. Hapus Barang")
        fmt.Println("5. Keluar")
        fmt.Print("Pilih menu: ")
        fmt.Scan(&pilihan)

        if pilihan == 1 {
            tambahBarang(&daftar)
        } else if pilihan == 2 {
            tampilkanData(&daftar)
        } else if pilihan == 3 {
            editBarang(&daftar)
        } else if pilihan == 4 {
            hapusBarang(&daftar)
        }
    }
}

func tambahBarang(T *TabBarang) {
    var b Barang
    if T.N < MaksBarang {
        fmt.Print("Nama barang: ")
        fmt.Scan(&b.Nama)
        fmt.Print("Jumlah stok: ")
        fmt.Scan(&b.Stok)

        T.Data[T.N] = b
        T.N = T.N + 1
    } else {
        fmt.Println("Kapasitas penuh.")
    }
}

func tampilkanData(T *TabBarang) {
    var urut int
    fmt.Println("Urut berdasarkan:")
    fmt.Println("1. Nama (Selection Sort)")
    fmt.Println("2. Stok (Insertion Sort)")
    fmt.Print("Pilih: ")
    fmt.Scan(&urut)

    if urut == 1 {
        selectionSortNama(T)
    } else if urut == 2 {
        insertionSortStok(T)
    }

    fmt.Println("=== Data Stok ===")
    for i := 0; i < T.N; i++ {
        fmt.Printf("%d. %s - %d\n", i+1, T.Data[i].Nama, T.Data[i].Stok)
    }
}

func editBarang(T *TabBarang) {
    var nama string
    var i int
    fmt.Print("Nama barang yang ingin diedit: ")
    fmt.Scan(&nama)

    i = sequentialSearch(*T, nama)
    if i != -1 {
        fmt.Print("Nama baru: ")
        fmt.Scan(&T.Data[i].Nama)
        fmt.Print("Stok baru: ")
        fmt.Scan(&T.Data[i].Stok)
    } else {
        fmt.Println("Barang tidak ditemukan.")
    }
}

func hapusBarang(T *TabBarang) {
    var nama string
    var i, j int
    fmt.Print("Nama barang yang ingin dihapus: ")
    fmt.Scan(&nama)

    i = sequentialSearch(*T, nama)
    if i != -1 {
        for j = i; j < T.N-1; j++ {
            T.Data[j] = T.Data[j+1]
        }
        T.N = T.N - 1
        fmt.Println("Barang dihapus.")
    } else {
        fmt.Println("Barang tidak ditemukan.")
    }
}

func sequentialSearch(T TabBarang, nama string) int {
    var i int
    for i = 0; i < T.N; i++ {
        if T.Data[i].Nama == nama {
            return i
        }
    }
    return -1
}

func selectionSortNama(T *TabBarang) {
    var i, j, idx int
    var tmp Barang
    for i = 0; i < T.N-1; i++ {
        idx = i
        for j = i + 1; j < T.N; j++ {
            if T.Data[j].Nama < T.Data[idx].Nama {
                idx = j
            }
        }
        tmp = T.Data[i]
        T.Data[i] = T.Data[idx]
        T.Data[idx] = tmp
    }
}

func insertionSortStok(T *TabBarang) {
    var i, j int
    var tmp Barang
    for i = 1; i < T.N; i++ {
        tmp = T.Data[i]
        j = i - 1
        for j >= 0 && T.Data[j].Stok > tmp.Stok {
            T.Data[j+1] = T.Data[j]
            j = j - 1
        }
        T.Data[j+1] = tmp
    }
}