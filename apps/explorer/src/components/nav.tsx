'use client';

import Image from 'next/image';
import { NavLink } from './nav-link';
import Link from 'next/link';
import { Fragment, useEffect, useState } from 'react';
import { usePathname } from 'next/navigation';

export const Nav = ({}) => {
	const [open, setOpen] = useState(false);
	const [navHeight, setNavHeight] = useState(0);

	const pathname = usePathname();

	useEffect(() => {
		if (open) {
			setOpen(false);
		}
	}, [pathname]);

	useEffect(() => {
		if (open) {
			document.body.style.overflow = 'hidden';
			return () => {
				document.body.style.overflow = 'scroll';
			};
		}
	}, [open]);

	const links = (
		<Fragment>
			<NavLink href='/'>Home</NavLink>
			<NavLink href='/chains'>Chains</NavLink>
			<NavLink href='/blocks'>Blocks</NavLink>
			<NavLink href='/txs'>Transactions</NavLink>
			<NavLink href='/vouchers'>Vouchers</NavLink>
			<NavLink href='/tokens'>Tokens</NavLink>
			<NavLink href='/nfts'>Carbon Credits</NavLink>
			<NavLink href='https://gesia.gitbook.io/gesia-project-kr'>Docs</NavLink>
		</Fragment>
	);
	return (
		<div className='w-full flex border-b border-b-[#EAECED] bg-white sticky top-0 z-20'>
			<nav
				className='mx-auto w-full max-w-[1440px] px-10 py-4 flex items-center max-lg:!px-4 max-lg:py-5'
				ref={(ref) => {
					if (ref) {
						setNavHeight(ref.clientHeight);
					}
				}}
			>
				<Link href={'/'} className='w-[185px] h-[26px] max-lg:!w-[140px] max-lg:!h-[20px] relative'>
					<Image src='/logo.svg' alt='Logo' fill />
				</Link>
				<div className='ml-auto'>
					<div className='hidden max-lg:flex'>
						<button
							className='w-[30px] h-[30px] relative'
							onClick={() => {
								setOpen((open) => !open);
							}}
						>
							<Image src='/menu.png' fill alt='Menu' />
						</button>
					</div>

					<div className='max-lg:hidden'>{links}</div>
				</div>
			</nav>

			{open && (
				<div
					className='hidden max-lg:flex fixed left-0 right-0 bottom-0 bg-[#00000033]'
					style={{
						top: navHeight,
					}}
					onClick={() => {
						setOpen(false);
					}}
				>
					<div className='bg-white flex flex-col w-full self-start'>{links}</div>
				</div>
			)}
		</div>
	);
};
