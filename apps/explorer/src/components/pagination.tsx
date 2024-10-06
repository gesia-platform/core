import Image from 'next/image';

export const Pagination = ({ offset, size, totalSize, onOffsetChange }: { offset: number; size: number; totalSize: number; onOffsetChange: (offset: number) => void }) => {
	const currentPage = Math.ceil(offset / size) + 1;
	const totalPage = Math.floor((totalSize ?? 0) / size) + 1;

	return (
		<div className='grid grid-flow-col auto-cols-auto gap-x-2'>
			<button
				className='h-[30px] px-[10px] border border-[#EAECED] text-[14px] text-[#1C1E20] rounded-[8px]'
				onClick={() => {
					onOffsetChange(0);
				}}
			>
				First
			</button>

			<button
				className='w-[30px] h-[30px] flex border border-[#EAECED] rounded-[8px]'
				disabled={currentPage === 1}
				onClick={() => {
					onOffsetChange(offset - size);
				}}
			>
				<Image className='m-auto' width={5} height={10} alt='Left' src={'/left.png'} />
			</button>

			<div className='h-[30px] border border-[#EAECED]  text-[14px] text-[#1C1E20] rounded-[8px] px-[10px] flex items-center'>
				{currentPage} of {totalPage}
			</div>

			<button
				disabled={currentPage === totalPage}
				className='w-[30px] h-[30px] flex border border-[#EAECED] rounded-[8px]'
				onClick={() => {
					onOffsetChange(offset + size);
				}}
			>
				<Image className='m-auto' width={5} height={10} alt='Right' src={'/right.png'} />
			</button>

			<button
				className='h-[30px] px-[10px] border border-[#EAECED] text-[14px] text-[#1C1E20] rounded-[8px]'
				onClick={() => {
					onOffsetChange(totalPage * size - size);
				}}
			>
				Last
			</button>
		</div>
	);
};
