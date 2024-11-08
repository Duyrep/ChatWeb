import Navigation from '@/components/navigation';
import './globals.css';

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {

  return (
    <html lang="en">
      <body className="p-2">
        <Navigation />
        <div className="my-14">{children}</div>
      </body>
    </html>
  );
}
