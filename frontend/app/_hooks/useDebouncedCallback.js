import { useCallback, useRef, useEffect } from 'react'

function useDebouncedCallback(callback, delay=200) {
    const timeoutRef = useRef(null)
    const callbackRef = useRef(callback)

    // Always keep the latest callback in the ref
    useEffect(() => {
        callbackRef.current = callback
    }, [callback])

    const debouncedFn = useCallback((...args) => {
        if (timeoutRef.current) {
            clearTimeout(timeoutRef.current)
        }

        timeoutRef.current = setTimeout(() => {
            if (callbackRef.current) {
                callbackRef.current(...args)
            }
        }, delay)
    }, [delay])


    return debouncedFn
}

export default useDebouncedCallback

