import './globals.css'

export const metadata = {
  title: '5minslearn',
  description: 'Sample NextJs App',
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en">
      <body>{children}</body>
    </html>
  )
}
