PROGRAM APOTEK C++
///*PROGRAM AKHIR SEMESTER*///
///*TEMA: APOTEK*///
///*ARRAY, STUCTUR, FUNGSI*///
#include <stdio.h>
#include <conio.h>
#include <iostream.h>
#include <iomanip.h>
korps()
 {
 cout<<“\t\tNUSA MANDIRI HEALTHY\n”;
 }
alamat()
 {
 cout<<”       Jl. penasihatguru.com kalimantan selatan\n”;
 cout<<”                 Cabang pelaihari\n”;
 }
trims()
 {
 cout<<”                   TERIMAKASIH\n”;
 cout<<”                semoga lekas sembuh\n”;
 cout<<”       info lengkap : www.penasihatguru.com\n”;
 cout<<”                 CC : xxxxxxxxxx\n”;
 }
line()
 {
 cout<<“___________________________________________________\n”;
 }
long jl(long byk, long hrg)
 {
 return (byk*hrg);
 }
long ttl(long tot, long hrgjl)
 {
 return (tot+hrgjl);
 }
long sisa( long tn, long tl)
 {
 return (tn-tl);
 }


main()
{
int x, y, a=0, b=0, c=0, hr, thn;
char kasir[30], bln[15];
long total, tunai, kembali;
struct
 {
 char kode[10], nama[20];
 long banyak, harga, jual;
 } produk[50];
up:
clrscr();
korps();
alamat();
cout<<endl;;
line();
cout<<endl;
cout<<“\tnama kasir\t: “;cin>>kasir;
cout<<endl;
cout<<“\ttanggal\t\t: “;cin>>hr;
cout<<“\tbulan\t\t: “;cin>>bln;
cout<<“\ttahun\t\t: “;cin>>thn;
cout<<endl;
line;
ne:
clrscr();
total=0;
a=a+1;
b=1;
c=1;
reply:
clrscr();
korps();
cout<<endl;
cout<<“nomor transaksi\t: No. 00″<<a<<endl;
cout<<“banyak beli jenis produk\t: “;cin>>y;
if (y<=10)
 goto next;
else
 goto reply;
next:
cout<<endl;
cout<<“kode produk\t: PCL, VIT, OBH, PST, BLS, MKP”<<endl;
cout<<“\t\t  MYA, TLKA, BTD, ALK”<<endl;
cout<<endl;
line();
for (x=1;x<=y;x++)
 {
   cout<<“jenis beli produk ke- “<<b++<<endl;
   cout<<“masukan kode produk\t\t: “;cin>>produk[x].kode;
   cout<<“jumlah beli produk\t\t: “;cin>>produk[x].banyak;
   line();
   if (!strcmp(produk[x].kode,”PCL”)||!strcmp(produk[x].kode,”pcl”))
    {
      strcpy(produk[x].nama,”Paracetamol”);
      produk[x].harga=2000;
      }
   else if (!strcmp(produk[x].kode,”VIT”)||!strcmp(produk[x].kode,”vit”))
    {
      strcpy(produk[x].nama,”Vitamin”);
      produk[x].harga=10000;
      }
   else if (!strcmp(produk[x].kode,”OBH”)||!strcmp(produk[x].kode,”obh”))
    {
      strcpy(produk[x].nama,”OBH batuk”);
      produk[x].harga=16000;
      }
   else if (!strcmp(produk[x].kode,”PST”)||!strcmp(produk[x].kode,”pst”))
    {
      strcpy(produk[x].nama,”Ponstan”);
      produk[x].harga=3000;
      }
   else if (!strcmp(produk[x].kode,”BLS”)||!strcmp(produk[x].kode,”bls”))
    {
      strcpy(produk[x].nama,”Balsem”);
      produk[x].harga=6000;
      }
   else if (!strcmp(produk[x].kode,”MKP”)||!strcmp(produk[x].kode,”mkp”))
    {
      strcpy(produk[x].nama,”Minyak kayu putih”);
      produk[x].harga=8000;
      }
   else if (!strcmp(produk[x].kode,”MYA”)||!strcmp(produk[x].kode,”mya”))
    {
      strcpy(produk[x].nama,”Minyak angin”);
      produk[x].harga=11000;
      }
   else if (!strcmp(produk[x].kode,”TLKA”)||!strcmp(produk[x].kode,”tlka”))
    {
      strcpy(produk[x].nama,”Tolak angin”);
      produk[x].harga=3000;
      }
   else if (!strcmp(produk[x].kode,”BTD”)||!strcmp(produk[x].kode,”btd”))
    {
      strcpy(produk[x].nama,”Betadine”);
      produk[x].harga=3000;
      }
   else if (!strcmp(produk[x].kode,”ALK”)||!strcmp(produk[x].kode,”alk”))
    {
      strcpy(produk[x].nama,”Alkohol 70%”);
      produk[x].harga=20000;
      }
   else
    {
      strcpy(produk[x].nama,”??????”);
      produk[x].harga=0;
  }
   }
clrscr();
korps();
alamat();
cout<<endl;
line();
cout<<bln<<“/”<<hr<<“/”<<setiosflags(ios::left)<<setw(22)<<thn;
cout<<setiosflags(ios::right)<<setw(8)<<kasir<<“/”<<“NMH00″<<a<<endl;
line();
cout<<“No. Nama\t\tharga\tbanyak\tharga”<<endl;
cout<<”    produk\t\tproduk\tbeli\tjual”<<endl;
cout<<endl;
for (x=1;x<=y;x++)
 {
   cout<<setiosflags(ios::left)<<setw(4)<<c++;
   cout<<setiosflags(ios::left)<<setw(20)<<produk[x].nama;
   cout<<setiosflags(ios::left)<<setw(8)<<produk[x].harga;
 cout<<setiosflags(ios::left)<<setw(8)<<produk[x].banyak;
   produk[x].jual=jl(produk[x].harga,produk[x].banyak);
 cout<<setiosflags(ios::right)<<setw(4)<<produk[x].jual<<endl;
   total=ttl(total,produk[x].jual);
   }
cout<endl;
line();
cout<<“\t\t\ttotal\t      : “<<total<<endl;
cout<<“\t\t\ttunai\t      : “;cin>>tunai;
kembali=sisa(tunai,total);
cout<<“\t\t\tkembali\t      : “<<kembali;
cout<<endl<<endl;
trims();
getch();
char lagi;
clrscr();
cout<<“tindak lanjut! pilih!\n”;
cout<<“N= transaksi baru\n”;
cout<<“R= restart program\n”;
cout<<“X= close\n”;
cout<<“==>>”;cin>>lagi;
if (lagi==’R’||lagi==’r’)
 {
   goto up;
   }
else if (lagi==’N’||lagi==’n’)
 {
   goto ne;
 }
else
 {
   goto ex;
   }
ex:
getch();
}