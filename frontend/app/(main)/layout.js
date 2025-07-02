import Header from './_components/header'
import Navigation from './_components/navigation'

export default function MainLayout({ children }) {

  return (
    <>
      <Header />
      <Navigation />
      {children}
    </>
  )
}
