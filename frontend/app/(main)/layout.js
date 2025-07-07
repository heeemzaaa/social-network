import Header from './_components/header'
import Navigation from './_components/navigation'


let notification = {
  type: "",
  message: "",
}

export default function MainLayout({ children }) {
  return (
    <>
      <Header />
      <Navigation />
      {children}
    </>
  )
}
