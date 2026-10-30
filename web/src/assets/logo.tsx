/*
Copyright (C) 2023-2026 QuantumNous

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as
published by the Free Software Foundation, either version 3 of the
License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program. If not, see <https://www.gnu.org/licenses/>.

For commercial licensing, please contact support@quantumnous.com
*/
import { type SVGProps } from 'react'

import { cn } from '@/lib/utils'

export function Logo({ className, ...props }: SVGProps<SVGSVGElement>) {
  return (
    <svg
      id='applerouter-logo'
      viewBox='0 0 24 24'
      xmlns='http://www.w3.org/2000/svg'
      height='24'
      width='24'
      fill='none'
      stroke='currentColor'
      strokeWidth='2'
      strokeLinecap='round'
      strokeLinejoin='round'
      className={cn('size-6', className)}
      {...props}
    >
      <title>Apple Router</title>
      <path d='M12 7.6c-1.6-1.1-4.9-1.3-6.2 1.4-1.2 2.5-.4 6.6 1.9 9.4 1.2 1.5 2.7 1.9 4.3 1.1 1.6.8 3.1.4 4.3-1.1 2.3-2.8 3.1-6.9 1.9-9.4-1.3-2.7-4.6-2.5-6.2-1.4z' />
      <path d='M12 7.6c-.1-1.6.4-2.9 1.6-3.8' />
      <path d='M9.6 16.2h1.2a2.2 2.2 0 0 0 2.2-2.2v-.6' />
    </svg>
  )
}
