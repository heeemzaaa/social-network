
export default function RootLayout({ children }) {
  console.log(children)

  return (
    <html lang="en">
      <body>
        {children}
      </body>
    </html>
  );
}
