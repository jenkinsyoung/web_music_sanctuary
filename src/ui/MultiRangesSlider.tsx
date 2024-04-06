import React, {useCallback, useEffect, useRef, useState}  from 'react'

const MultiRangesSlider = ({min, max} : any) => {
  const [minVal, setMin] = useState(min);
  const [maxVal, setMax] = useState(max);
  const minValRef = useRef<any>()
  const maxValRef = useRef<any>()
  const rangeRef = useRef<any>()

  const getPercent = useCallback(
    (value : any) => Math.round(((value - min) / (max - min)) * 100),
    [min, max]
  )

  useEffect(()=>{
    if(maxValRef.current){
      const minPercent = getPercent(minVal)
      const maxPercent = getPercent(+maxValRef.current.value)

      if(rangeRef.current){
        rangeRef.current.style.left = `${minPercent}%`
        rangeRef.current.style.width = `${maxPercent - minPercent}%`
      }
    }
  }, [minVal, getPercent])

  useEffect(()=>{
    if(minValRef.current){
      const maxPercent = getPercent(+minValRef.current.value)
      const minPercent = getPercent(maxVal)

      if(rangeRef.current){
        rangeRef.current.style.width = `${maxPercent - minPercent}%`
      }
    }
  }, [maxVal, getPercent])

  return (
    <div className='flex justify-center items-center'>
        <input 
          type='range' 
          min={min}
          max={max}
          value={minVal}
          ref ={minValRef}
          onChange={(e)=>{
            const value = Math.min(+e.target.value, maxVal -1)
            setMin(value)
            e.target.value = value.toString()
          }}
          className='thumb z-[3]'
        />
        <input 
          type='range' 
          min={min}
          max={max}
          ref ={maxValRef}
          value={maxVal}
          onChange={(e)=>{
            const value = Math.max(+e.target.value, minVal +1)
            setMax(value)
            e.target.value = value.toString()
          }}
          className='thumb z-[4]'
        />
    
        <div className='relative w-52'>
             <div className='flex' style={{marginTop: '12px', width: '390px', justifyContent: 'space-between'}}>  
            <div className='absolute left-2 text-xs mt-14'>{minVal}</div>
            <div className='absolute right-2 text-xs mt-14'>{maxVal}</div>
            </div> 
        </div>
    
    </div>
  )
}

export default MultiRangesSlider