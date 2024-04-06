import MultiRangesSlider from '@/ui/MultiRangesSlider';
const PriceFilter = ({min, max} : any) => {
  return(
    <div className='h-40 p-3 space-y-2'>
      <h2 style={{fontSize: '18px', fontWeight: '600'}}>Цена:</h2>
      <MultiRangesSlider
        min={min}
        max={max}
         />
    </div>
  )
};

export default PriceFilter;