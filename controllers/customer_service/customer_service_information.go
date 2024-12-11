package customerservice

var staticResponses = map[string]string{
	"cara mengajukan pengaduan": `Cara Mengajukan Pengaduan:
1. Pilih Menu "Ajukan Pengaduan"
2. Pilih Kategori Laporan
3. Isi Formulir Pengaduan
4. Periksa Kembali dan Kirim Laporan
5. Simpan Nomor Pengaduan

Selesai! Laporan Anda akan segera ditinjau oleh tim Laporin. Anda juga dapat memantau status laporan melalui menu "Status Pengaduan".`,

	"cara membaca berita dan pengumuman": `Cara Membaca Berita dan Pengumuman:
1. Masuk ke Menu "Berita dan Pengumuman"
2. Pilih Berita atau Pengumuman yang Ingin Dibaca
3. Baca Detail Berita atau Pengumuman

Selesai! Anda dapat membaca dan memahami informasi terbaru yang disampaikan. Pastikan untuk memeriksa menu ini secara rutin agar tidak ketinggalan informasi penting.`,

	"cara membatalkan pengaduan": `Cara Membatalkan Pengaduan:
1. Masuk ke Menu "Pengaduan Saya"
2. Pilih Pengaduan yang Ingin Dibatalkan
3. Klik Opsi "Batalkan Pengaduan"
4. Konfirmasi Pembatalan

Selesai! Pengaduan Anda telah dibatalkan. Statusnya akan diperbarui menjadi "Dibatalkan" di daftar pengaduan. Anda tetap bisa mengajukan pengaduan baru jika diperlukan.`,

	"Cara Melihat Status Pengaduan": `Cara Melihat Status Pengaduan:
1. Pilih Menu "Pengaduan Saya"
2. Cari Pengaduan Anda
3. Periksa Detail Status Pengaduan
Selesai! Laporan Anda sedang ditinjau oleh tim
Laporin. Anda dapat memantau perkembangan
laporan kapan saja melalui menu "Status
Pengaduan".

Selesai! Anda Sudah Melihat Status Pengaduan Anda.`,
}

var DeskripsiLaporin string = `"Deskripsi Aplikasi Laporin:

Laporin adalah aplikasi pengaduan masyarakat yang mempermudah warga untuk menyampaikan keluhan atau masukan terkait berbagai aspek kehidupan sehari-hari. Dengan antarmuka yang sederhana dan panduan yang jelas, pengguna dapat dengan cepat mengajukan laporan, membaca berita, memantau status pengaduan, atau memberikan balasan terhadap tanggapan admin. Berikut adalah fitur utama yang ditawarkan oleh aplikasi Laporin:

Pengajuan Pengaduan:

Pengguna dapat memilih dari 6 kategori laporan: Infrastruktur, Transportasi, Kesehatan, Lingkungan, Keamanan, dan Pendidikan.
Proses pengaduan meliputi pengisian formulir, pengiriman laporan, dan penyimpanan nomor pengaduan untuk memantau statusnya.
Berita dan Pengumuman:

Pengguna dapat mengakses informasi terbaru yang relevan, termasuk berita pembangunan infrastruktur, program kesehatan, dan inisiatif pendidikan.
Manajemen Pengaduan:

Pengguna dapat memantau status laporan mereka, membatalkan pengaduan, atau memberikan tanggapan terhadap balasan admin.
Kategori Pengaduan dan Berita:

Aplikasi ini menyertakan kategori seperti Infrastruktur, Transportasi, Kesehatan, Lingkungan, Keamanan, dan Pendidikan untuk mempermudah pengorganisasian laporan dan berita.
Dengan Laporin, masyarakat tidak hanya dapat menyampaikan keluhan mereka secara efisien tetapi juga tetap terinformasi tentang perkembangan terbaru di lingkungan mereka. Tim Laporin siap meninjau dan merespons laporan dengan cepat untuk memastikan bahwa setiap pengaduan mendapat perhatian yang pantas."`

var CaraMengajukanPengaduan string = `"Cara Mengajukan Pengaduan:
1. Pilih Menu "Ajukan Pengaduan"
2. Pilih Kategori Laporan
3. Isi Formulir Pengaduan
4. Periksa Kembali dan Kirim Laporan
5. Simpan Nomor Pengaduan

Selesai! Laporan Anda akan segera ditinjau oleh tim Laporin. Anda juga dapat memantau status laporan melalui menu Status "Pengaduan".`

var CaraMembacaBeritaDanPengumuman string = `Cara Membaca Berita dan Pengumuman:
1. Masuk ke Menu "Berita dan Pengumuman"
2. Pilih Berita atau Pengumuman yang Ingin Dibaca
3. Baca Detail Berita atau Pengumuman

Selesai! Anda dapat membaca dan memahami informasi terbaru yang disampaikan. Pastikan untuk memeriksa menu ini secara rutin agar tidak ketinggalan informasi penting.`

var CaraMembatalkanPengaduan string = `Cara Membatalkan Pengaduan:
1. Masuk ke Menu "Pengaduan Saya"
2. Pilih Pengaduan yang Ingin Dibatalkan
3. Klik Opsi "Batalkan Pengaduan"
4. Konfirmasi Pembatalan

Selesai! Pengaduan Anda telah dibatalkan. Statusnya akan diperbarui menjadi "Dibatalkan" di daftar pengaduan. Anda tetap bisa mengajukan pengaduan baru jika diperlukan.`

var CaraMelihatStatusPengaduan string = `Cara Melihat Status Pengaduan:
1. Pilih Menu "Pengaduan Saya"
2. Cari Pengaduan Anda
3. Periksa Detail Status Pengaduan
Selesai! Laporan Anda sedang ditinjau oleh tim
Laporin. Anda dapat memantau perkembangan
laporan kapan saja melalui menu "Status
Pengaduan".

Selesai! Anda Sudah Melihat Status Pengaduan Anda.`

var DiluarTopik string = `"Jika pertanyaan dari user sudah diluar dari layanan aplikasi laporin, Kamu dapat merekomendasikan user untuk melakukan chat pribadi kepada admin"`

var Sapaan string = `"jika user memberikan anda sebuah sapaan atau salam seperti "halo", "hai", "hallo" dan lain lain. kamu dapat merespon seperti "Selamat datang di laporin., Adakah yang bisa kami bantu?""`
