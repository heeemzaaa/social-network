import { useEffect, useState } from "react";





const useDebounce = (value, delay = 500) => {
    const [debouncedValue, setDebouncedValue] = useState(value)
    useEffect(() => {
        const timeout = setTimeout(() => {
            setDebouncedValue(value)
        }, delay)
        // ila tchangeat l value 9bl maysali timeout
        return () => clearTimeout(timeout)
    }, [value, delay])
    return debouncedValue
}


export default useDebounce