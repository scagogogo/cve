#!/usr/bin/env python3
"""Generate a feature mind map for the CVE Utils project (no emoji, clean style)."""

import matplotlib
matplotlib.use('Agg')
import matplotlib.pyplot as plt
from matplotlib.patches import FancyBboxPatch, FancyArrowPatch
import numpy as np

# Pastel color palette
PASTEL = {
    'root': '#2d3436',
    'format': '#ff6b6b',
    'extract': '#4ecdc4',
    'compare': '#45b7d1',
    'filter': '#96ceb4',
    'generate': '#f9ca24',
    'set_ops': '#a29bfe',
    'batch': '#fd79a8',
    'range': '#fab1a0',
    'stats': '#00cec9',
}

def draw_mindmap():
    fig, ax = plt.subplots(1, 1, figsize=(26, 16))
    ax.set_xlim(-13, 13)
    ax.set_ylim(-8, 8)
    ax.set_aspect('equal')
    ax.axis('off')
    fig.patch.set_facecolor('#fafbfc')

    # Root node
    root_x, root_y = 0, 0
    root_box = FancyBboxPatch((root_x - 2.8, root_y - 0.55), 5.6, 1.1,
                               boxstyle="round,pad=0.2", facecolor='#2d3436',
                               edgecolor='#636e72', linewidth=3, zorder=10)
    ax.add_patch(root_box)
    ax.text(root_x, root_y + 0.08, 'CVE Utils', fontsize=24, fontweight='bold',
            color='white', ha='center', va='center', zorder=11, fontfamily='sans-serif')
    ax.text(root_x, root_y - 0.35, '30+ Functions | 9 Modules', fontsize=10,
            color='#b2bec3', ha='center', va='center', zorder=11, fontfamily='sans-serif')

    # Branch definitions: (angle, distance, label, color_key, children)
    branches = [
        (140, 4.5, 'Format &\nValidation', 'format', [
            'Format()', 'IsCve()', 'IsContainsCve()',
            'ValidateCve()', 'IsCveYearOk()',
            'IsCveYearOkWithCutoff()'
        ]),
        (100, 4.5, 'Extraction', 'extract', [
            'ExtractCve()', 'ExtractFirstCve()',
            'ExtractLastCve()', 'Split()',
            'ExtractCveYear()', 'ExtractCveSeq()',
            'ExtractCveYearAsInt()', 'ExtractCveSeqAsInt()'
        ]),
        (60, 4.5, 'Compare &\nSorting', 'compare', [
            'CompareCves()', 'CompareByYear()',
            'SubByYear()', 'SortCves()'
        ]),
        (20, 4.5, 'Filter &\nGrouping', 'filter', [
            'FilterCvesByYear()', 'FilterCvesByYearRange()',
            'GetRecentCves()', 'GroupByYear()',
            'RemoveDuplicateCves()'
        ]),
        (-20, 4.5, 'Generation', 'generate', [
            'GenerateCve()', 'GenerateFakeCve()',
            'FormatSeq()'
        ]),
        (-60, 4.5, 'Set\nOperations', 'set_ops', [
            'IntersectCves()', 'UnionCves()',
            'DiffCves()'
        ]),
        (-100, 4.5, 'Batch\nValidation', 'batch', [
            'ValidateCves()', 'FilterValidCves()'
        ]),
        (-140, 4.5, 'Range &\nPattern', 'range', [
            'ParseCveRange()', 'IsCvesConsecutive()',
            'FilterCvesByPattern()'
        ]),
        (180, 4.5, 'Statistical\nAnalysis', 'stats', [
            'CountByYear()', 'YearRange()', 'SeqRange()'
        ]),
    ]

    for angle_deg, dist, label, color_key, children in branches:
        angle = np.radians(angle_deg)
        branch_x = root_x + dist * np.cos(angle)
        branch_y = root_y + dist * np.sin(angle)
        color = PASTEL[color_key]

        # Curved connection line
        ctrl_offset = 1.0
        ctrl_x = root_x + dist * 0.55 * np.cos(angle) + ctrl_offset * np.sin(angle)
        ctrl_y = root_y + dist * 0.55 * np.sin(angle) - ctrl_offset * np.cos(angle)

        t_vals = np.linspace(0, 1, 60)
        curve_x = (1-t_vals)**2 * root_x + 2*(1-t_vals)*t_vals * ctrl_x + t_vals**2 * branch_x
        curve_y = (1-t_vals)**2 * root_y + 2*(1-t_vals)*t_vals * ctrl_y + t_vals**2 * branch_y
        ax.plot(curve_x, curve_y, color=color, linewidth=3.5, alpha=0.5, zorder=2, solid_capstyle='round')

        # Branch node
        is_two_line = '\n' in label
        bbox_h = 0.75 if not is_two_line else 0.95
        bbox_w = 1.9
        branch_box = FancyBboxPatch(
            (branch_x - bbox_w/2, branch_y - bbox_h/2), bbox_w, bbox_h,
            boxstyle="round,pad=0.12", facecolor=color,
            edgecolor='white', linewidth=2, zorder=8, alpha=0.92
        )
        ax.add_patch(branch_box)
        text_color = 'white' if color_key != 'generate' else '#2d3436'
        ax.text(branch_x, branch_y, label, fontsize=11, fontweight='bold',
                color=text_color, ha='center', va='center', zorder=9,
                fontfamily='sans-serif', linespacing=1.1)

        # Children
        n_children = len(children)
        fan_half = 28 if angle_deg not in [180, 0] else 35
        child_start_angle = angle_deg - fan_half
        child_end_angle = angle_deg + fan_half
        child_dist = 2.8

        for i, child in enumerate(children):
            if n_children > 1:
                child_angle = np.radians(child_start_angle + (child_end_angle - child_start_angle) * i / (n_children - 1))
            else:
                child_angle = angle

            child_x = branch_x + child_dist * np.cos(child_angle)
            child_y = branch_y + child_dist * np.sin(child_angle)

            # Connection to child
            ax.plot([branch_x, child_x], [branch_y, child_y],
                    color=color, linewidth=1.5, alpha=0.35, zorder=1)

            # Child node
            text_w = max(len(child) * 0.075, 0.7)
            child_box = FancyBboxPatch(
                (child_x - text_w - 0.12, child_y - 0.22), (text_w + 0.12) * 2, 0.44,
                boxstyle="round,pad=0.06", facecolor='white',
                edgecolor=color, linewidth=1.5, zorder=6, alpha=0.95
            )
            ax.add_patch(child_box)
            ax.text(child_x, child_y, child, fontsize=8, color='#2d3436',
                    ha='center', va='center', zorder=7, fontfamily='monospace',
                    fontweight='medium')

    # Title
    ax.text(0, 7.5, 'CVE Utils - Feature Map', fontsize=28, fontweight='bold',
            color='#2d3436', ha='center', va='center', fontfamily='sans-serif')
    ax.text(0, 6.8, 'Comprehensive Go library for CVE identifier processing',
            fontsize=13, color='#636e72', ha='center', va='center', style='italic')

    plt.tight_layout(pad=0.5)
    plt.savefig('/home/cc11001100/github/scagogogo/cve-skills/docs/images/feature-map.png',
                dpi=150, bbox_inches='tight', facecolor='#fafbfc')
    print("Feature map saved!")


if __name__ == '__main__':
    draw_mindmap()
