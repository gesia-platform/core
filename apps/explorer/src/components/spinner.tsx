import React from 'react';
import { SyncLoader } from 'react-spinners';

export const Spinner = ({ height = '200px', color = '#000000', size = 10 }) => {
	return (
		<div className={`flex justify-center items-center`} style={{ minHeight: height }}>
			<SyncLoader size={size} color={color} />
		</div>
	);
};
