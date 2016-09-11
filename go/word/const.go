package main

const (
	OutputFile  = `output.doc`
	HeaderFile  = "word_header"
	FooterFile  = "word_footer"
	TableHeader = `
		<w:tbl>
				<w:tblPr>
					<w:tblW w:w="0" w:type="auto"/>
					<w:tblBorders>
						<w:top w:val="thick-thin-medium-gap" w:sz="24" wx:bdrwidth="120" w:space="0" w:color="A5A5A5"/>
					</w:tblBorders>
					<w:tblLook w:val="04A0"/>
				</w:tblPr>
				<w:tblGrid>
					<w:gridCol w:w="2840"/>
					<w:gridCol w:w="2841"/>
					<w:gridCol w:w="2841"/>
				</w:tblGrid>`
	TableFooter = `
		</w:tbl>
	`
	TableRaw = `
	<w:tr wsp:rsidR="00AF5A68" wsp:rsidTr="00AF5A68">
					<w:tc>
						<w:tcPr>
							<w:tcW w:w="2840" w:type="dxa"/>
							<w:shd w:val="clear" w:color="auto" w:fill="auto"/>
						</w:tcPr>
						<w:p wsp:rsidR="00AF5A68" wsp:rsidRDefault="00AF5A68" wsp:rsidP="00AF5A68">
							<w:pPr>
								<w:tabs>
									<w:tab w:val="center" w:pos="1312"/>
								</w:tabs>
							</w:pPr>
							<w:r>
								<w:t>%s</w:t>
							</w:r>
						</w:p>
					</w:tc>
					<w:tc>
						<w:tcPr>
							<w:tcW w:w="2841" w:type="dxa"/>
							<w:shd w:val="clear" w:color="auto" w:fill="auto"/>
						</w:tcPr>
						<w:p wsp:rsidR="00AF5A68" wsp:rsidRDefault="00AF5A68">
							<w:r>
								<w:t>%s</w:t>
							</w:r>
						</w:p>
					</w:tc>
					<w:tc>
						<w:tcPr>
							<w:tcW w:w="2841" w:type="dxa"/>
							<w:shd w:val="clear" w:color="auto" w:fill="auto"/>
						</w:tcPr>
						<w:p wsp:rsidR="00AF5A68" wsp:rsidRDefault="00AF5A68">
							<w:r>
								<w:t>%s</w:t>
							</w:r>
						</w:p>
					</w:tc>
				</w:tr>
	`
)
