import MultiRangesSlider from '@/ui/MultiRangesSlider';
const PriceFilter = () => {
  return(
    <div className='h-40 p-5 rounded-md space-y-2'>
      <h2 style={{fontSize: '18px', fontWeight: '600'}}>Цена:</h2>
      <MultiRangesSlider
        min={0}
        max={100}
         />
    </div>
  )
};

export default PriceFilter;